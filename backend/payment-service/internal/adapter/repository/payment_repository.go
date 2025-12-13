package repository

import (
	"context"
	"payment-service/internal/core/domain/entity"
	"payment-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PaymentRepositoryInterface interface {
	CreatePayment(ctx context.Context, payment entity.PaymentEntity) error
	LogPayment(ctx context.Context, paymentID uint, status string) error
	UpdateStatusByOrderCode(ctx context.Context, orderID uint, status string) error
}

type paymentRepository struct {
	db *gorm.DB
}

// UpdateStatusByOrderCode implements PaymentRepositoryInterface.
func (p *paymentRepository) UpdateStatusByOrderCode(ctx context.Context, orderID uint, status string) error {
	modelPayment := model.Payment{}

	if err := p.db.Where("order_id = ?", orderID).First(&modelPayment).Error; err != nil {
		log.Errorf("[PaymentRepository] UpdateStatusByOrderCode-1: %v", err)
		return err
	}

	modelPayment.PaymentStatus = status

	if err := p.db.Save(&modelPayment).Error; err != nil {
		log.Errorf("[PaymentRepository] UpdateStatusByOrderCode-2: %v", err)
		return err
	}

	return nil
}

// LogPayment implements PaymentRepositoryInterface.
func (p *paymentRepository) LogPayment(ctx context.Context, paymentID uint, status string) error {
	logPayment := model.PaymentLog{
		PaymentID: paymentID,
		Status:    status,
	}

	if err := p.db.Create(&logPayment).Error; err != nil {
		log.Errorf("[PaymentRepository] LogPayment-1: %v", err)
		return err
	}

	return nil
}

// CreatePayment implements PaymentRepositoryInterface.
func (p *paymentRepository) CreatePayment(ctx context.Context, payment entity.PaymentEntity) error {
	modelPayment := model.Payment{
		OrderID:          payment.OrderID,
		UserID:           payment.UserID,
		PaymentMethod:    payment.PaymentMethod,
		PaymentStatus:    payment.PaymentStatus,
		PaymentGatewayID: &payment.PaymentGatewayID,
		GrossAmount:      payment.GrossAmount,
		PaymentURL:       &payment.PaymentURL,
	}

	if err := p.db.Create(&modelPayment).Error; err != nil {
		log.Errorf("[PaymentRepository] Create-1: %v", err)
		return err
	}

	return p.LogPayment(ctx, modelPayment.ID, modelPayment.PaymentStatus)
}

func NewPaymentRepository(db *gorm.DB) PaymentRepositoryInterface {
	return &paymentRepository{db: db}
}
