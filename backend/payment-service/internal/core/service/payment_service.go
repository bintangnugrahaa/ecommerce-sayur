package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"payment-service/config"
	httpclient "payment-service/internal/adapter/http_client"
	"payment-service/internal/adapter/message"
	"payment-service/internal/adapter/repository"
	"payment-service/internal/core/domain/entity"
	"strconv"

	"github.com/labstack/gommon/log"
)

type PaymentServiceInterface interface {
	ProcessPayment(ctx context.Context, payment entity.PaymentEntity, accessToken string) (*entity.PaymentEntity, error)
	UpdateStatusByOrderCode(ctx context.Context, orderCode, status, accessToken string) error
	GetAll(ctx context.Context, req entity.PaymentQueryStringRequest, accessToken string) ([]entity.PaymentEntity, int64, int64, error)
}

type paymentService struct {
	repo                repository.PaymentRepositoryInterface
	httpClientToService httpclient.HttpClientToService
	midtrans            httpclient.MidtransClientInterface
	cfg                 *config.Config
	publisherRabbitMQ   message.PublishRabbitMQInterface
}

// GetAll implements [PaymentServiceInterface].
func (p *paymentService) GetAll(ctx context.Context, req entity.PaymentQueryStringRequest, accessToken string) ([]entity.PaymentEntity, int64, int64, error) {
	results, count, total, err := p.repo.GetAll(ctx, req)
	if err != nil {
		log.Errorf("[PaymentService] GetAll-1: %v", err)
		return nil, 0, 0, err
	}

	var token map[string]interface{}
	err = json.Unmarshal([]byte(accessToken), &token)
	if err != nil {
		log.Errorf("[PaymentService] GetAll-2: %v", err)
		return nil, 0, 0, err
	}
	for key, val := range results {
		orderDetail, err := p.httpClientOrderService(int64(val.OrderID), token["token"].(string))
		if err != nil {
			log.Errorf("[PaymentService] GetAll-3: %v", err)
			return nil, 0, 0, err
		}
		results[key].OrderCode = orderDetail.OrderCode
		results[key].OrderShippingType = orderDetail.ShippingType
	}

	return results, count, total, nil
}

// UpdateStatusByOrderCode implements PaymentServiceInterface.
func (p *paymentService) UpdateStatusByOrderCode(ctx context.Context, orderCode string, status string, accessToken string) error {
	orderDetail, err := p.httpClientOrderByCodeService(orderCode, accessToken)
	if err != nil {
		log.Errorf("[PaymentService] UpdateStatusByOrderCode-1: %v", err)
		return err
	}

	if err = p.repo.UpdateStatusByOrderCode(ctx, uint(orderDetail.ID), status); err != nil {
		log.Errorf("[PaymentService] UpdateStatusByOrderCode-2: %v", err)
		return err
	}

	return nil
}

// ProcessPayment implements PaymentServiceInterface.
func (p *paymentService) ProcessPayment(ctx context.Context, payment entity.PaymentEntity, accessToken string) (*entity.PaymentEntity, error) {
	if payment.PaymentMethod == "cod" {
		payment.PaymentStatus = "Success"

		if err := p.repo.CreatePayment(ctx, payment); err != nil {
			log.Errorf("[PaymentService] ProcessPayment-1: %v", err)
			return nil, err
		}

		if err := p.publisherRabbitMQ.PublishPaymentSuccess(payment); err != nil {
			log.Errorf("[PaymentService] ProcessPayment-2: %v", err)
		}

		return &payment, nil
	}

	if payment.PaymentMethod == "midtrans" {
		var token map[string]interface{}
		err := json.Unmarshal([]byte(accessToken), &token)
		if err != nil {
			log.Errorf("[PaymentService] ProcessPayment-3: %v", err)
			return nil, err
		}

		userResponse, err := p.httpClientUserService(token["token"].(string))
		if err != nil {
			log.Errorf("[PaymentService] ProcessPayment-4: %v", err)
			return nil, err
		}

		orderDetail, err := p.httpClientOrderService(int64(payment.OrderID), token["token"].(string))
		if err != nil {
			log.Errorf("[PaymentService] ProcessPayment-5: %v", err)
			return nil, err
		}

		transactionID, err := p.midtrans.CreateTransaction(orderDetail.OrderCode, int64(payment.GrossAmount), userResponse.Name, userResponse.Email)
		if err != nil {
			log.Errorf("[PaymentService] ProcessPayment-6: %v", err)
			return nil, err
		}
		payment.PaymentStatus = "Pending"
		payment.PaymentGatewayID = transactionID

		if err := p.repo.CreatePayment(ctx, payment); err != nil {
			log.Errorf("[PaymentService] ProcessPayment-7: %v", err)
			return nil, err
		}

		if err := p.publisherRabbitMQ.PublishPaymentSuccess(payment); err != nil {
			log.Errorf("[PaymentService] ProcessPayment-8: %v", err)
		}

		return &payment, nil
	}

	return nil, errors.New("Invalid payment method")
}

func (p *paymentService) httpClientOrderService(orderId int64, accessToken string) (*entity.OrderDetailHttpResponse, error) {
	baseUrlOrder := fmt.Sprintf("%s/%s", p.cfg.App.OrderServiceUrl, "auth/orders/"+strconv.FormatInt(orderId, 10))
	header := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Accept":        "application/json",
	}
	dataOrder, err := p.httpClientToService.CallURL("GET", baseUrlOrder, header, nil)
	if err != nil {
		log.Errorf("[PaymentService] httpClientOrderService-1: %v", err)
		return nil, err
	}

	defer dataOrder.Body.Close()

	body, err := io.ReadAll(dataOrder.Body)
	if err != nil {
		log.Errorf("[PaymentService] httpClientOrderService-2: %v", err)
		return nil, err
	}

	var orderDetail entity.OrderHttpClientResponse
	err = json.Unmarshal([]byte(body), &orderDetail)
	if err != nil {
		log.Errorf("[PaymentService] httpClientOrderService-3: %v", err)
		return nil, err
	}

	return &orderDetail.Data, nil
}

func (p *paymentService) httpClientUserService(accessToken string) (*entity.ProfileHttpResponse, error) {
	baseUrlUser := fmt.Sprintf("%s/%s", p.cfg.App.UserServiceUrl, "auth/profile")
	header := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Accept":        "application/json",
	}
	dataUser, err := p.httpClientToService.CallURL("GET", baseUrlUser, header, nil)
	if err != nil {
		log.Errorf("[PaymentService] httpClientUserService-1: %v", err)
		return nil, err
	}

	defer dataUser.Body.Close()

	body, err := io.ReadAll(dataUser.Body)
	if err != nil {
		log.Errorf("[PaymentService] httpClientUserService-2: %v", err)
		return nil, err
	}

	var userResponse entity.UserHttpClientResponse
	err = json.Unmarshal([]byte(body), &userResponse)
	if err != nil {
		log.Errorf("[PaymentService] httpClientUserService-3: %v", err)
		return nil, err
	}

	return &userResponse.Data, nil
}

func (p *paymentService) httpClientOrderByCodeService(orderCode string, accessToken string) (*entity.OrderDetailHttpResponse, error) {
	baseUrlOrder := fmt.Sprintf("%s/%s", p.cfg.App.OrderServiceUrl, "auth/orders/"+orderCode+"/code")
	header := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Accept":        "application/json",
	}
	dataOrder, err := p.httpClientToService.CallURL("GET", baseUrlOrder, header, nil)
	if err != nil {
		log.Errorf("[PaymentService] httpClientOrderByCodeService-1: %v", err)
		return nil, err
	}

	defer dataOrder.Body.Close()

	body, err := io.ReadAll(dataOrder.Body)
	if err != nil {
		log.Errorf("[PaymentService] httpClientOrderByCodeService-2: %v", err)
		return nil, err
	}

	var orderDetail entity.OrderHttpClientResponse
	err = json.Unmarshal([]byte(body), &orderDetail)
	if err != nil {
		log.Errorf("[PaymentService] httpClientOrderByCodeService-3: %v", err)
		return nil, err
	}

	return &orderDetail.Data, nil
}

func NewPaymentService(repo repository.PaymentRepositoryInterface, cfg *config.Config, httpClientToService httpclient.HttpClientToService, midtrans httpclient.MidtransClientInterface, publisherRabbitMQ message.PublishRabbitMQInterface) PaymentServiceInterface {
	return &paymentService{
		repo:                repo,
		httpClientToService: httpClientToService,
		midtrans:            midtrans,
		cfg:                 cfg,
		publisherRabbitMQ:   publisherRabbitMQ,
	}
}
