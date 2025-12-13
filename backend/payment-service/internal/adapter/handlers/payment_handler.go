package handlers

import (
	"net/http"
	"payment-service/config"
	"payment-service/internal/adapter"
	"payment-service/internal/adapter/handlers/request"
	"payment-service/internal/adapter/handlers/response"
	"payment-service/internal/core/domain/entity"
	"payment-service/internal/core/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type PaymentHandlerInterface interface {
	Create(c echo.Context) error
}

type paymentHandler struct {
	paymentService service.PaymentServiceInterface
}

// Create implements PaymentHandlerInterface.
func (p *paymentHandler) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = request.PaymentRequest{}
	)

	user := c.Get("user").(string)
	if user == "" {
		log.Errorf("[PaymentHandler-1] Create: %s", "data token not found")
		return c.JSON(http.StatusUnauthorized, response.ResponseDefault("data token not found", nil))
	}

	if err := c.Bind(&req); err != nil {
		log.Errorf("[PaymentHandler-2] Create: %v", err)
		return c.JSON(http.StatusBadRequest, response.ResponseDefault(err.Error(), nil))
	}

	if err := c.Validate(&req); err != nil {
		log.Errorf("[PaymentHandler-3] Create: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, response.ResponseDefault(err.Error(), nil))
	}

	paymentEntity := entity.PaymentEntity{
		OrderID:       req.OrderID,
		PaymentMethod: req.PaymentMethod,
		GrossAmount:   float64(req.GrossAmount),
		UserID:        req.UserID,
		Remarks:       req.Remarks,
	}

	result, err := p.paymentService.ProcessPayment(ctx, paymentEntity, user)
	if err != nil {
		log.Errorf("[PaymentHandler-4] Create: %v", err)
		return c.JSON(http.StatusInternalServerError, response.ResponseDefault(err.Error(), nil))
	}

	responPayment := map[string]interface{}{
		"payment_token": result.PaymentGatewayID,
	}

	return c.JSON(http.StatusCreated, response.ResponseDefault("success", responPayment))
}

func NewPaymentHandler(paymentService service.PaymentServiceInterface, e *echo.Echo, cfg *config.Config) PaymentHandlerInterface {
	paymentHandler := &paymentHandler{
		paymentService: paymentService,
	}
	e.Use(middleware.Recover())
	mid := adapter.NewMiddlewareAdapter(cfg)
	authGroup := e.Group("auth", mid.CheckToken())
	authGroup.POST("/payments", paymentHandler.Create)
	return paymentHandler
}
