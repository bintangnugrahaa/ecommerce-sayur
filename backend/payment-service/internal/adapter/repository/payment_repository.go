package repository

import (
	"context"
	"payment-service/internal/core/domain/entity"
	"payment-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PaymentRepositoryInterface interface {
	CreatePayment(ctx context.Context, payment *entity.PaymentEntity) error
}

type paymentRepository struct {
	db *gorm.DB
}

// CreatePayment implements PaymentRepositoryInterface.
func (p *paymentRepository) CreatePayment(ctx context.Context, payment *entity.PaymentEntity) error {
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

	return nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepositoryInterface {
	return &paymentRepository{db: db}
}
