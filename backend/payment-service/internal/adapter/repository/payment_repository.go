package repository

import (
	"context"
	"errors"
	"math"
	"payment-service/internal/core/domain/entity"
	"payment-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PaymentRepositoryInterface interface {
	CreatePayment(ctx context.Context, payment entity.PaymentEntity) error
	LogPayment(ctx context.Context, paymentID uint, status string) error
	UpdateStatusByOrderCode(ctx context.Context, orderID uint, status string) error
	GetAll(ctx context.Context, req entity.PaymentQueryStringRequest) ([]entity.PaymentEntity, int64, int64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

// GetAll implements [PaymentRepositoryInterface].
func (p *paymentRepository) GetAll(ctx context.Context, req entity.PaymentQueryStringRequest) ([]entity.PaymentEntity, int64, int64, error) {
	modelPayments := []model.Payment{}
	var countData int64
	offset := (req.Page - 1) * req.Limit

	sqlMain := p.db.
		Where("order_code ILIKE ? OR payment_method ILIKE ? OR status ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Status+"%")

	if req.UserID != 0 {
		sqlMain = sqlMain.Where("user_id = ?", req.UserID)
	}

	if err := sqlMain.Model(&modelPayments).Count(&countData).Error; err != nil {
		log.Errorf("[PaymentRepository-1] GetAll: %v", err)
		return nil, 0, 0, err
	}

	totalPage := int(math.Ceil(float64(countData) / float64(req.Limit)))
	if err := sqlMain.Order("created_at DESC").Limit(int(req.Limit)).Offset(int(offset)).Find(&modelPayments).Error; err != nil {
		log.Errorf("[PaymentRepository-2] GetAll: %v", err)
		return nil, 0, 0, err
	}

	if len(modelPayments) == 0 {
		err := errors.New("404")
		log.Infof("[PaymentRepository-3] GetAll: No payment found")
		return nil, 0, 0, err
	}

	entities := []entity.PaymentEntity{}
	for _, val := range modelPayments {
		entities = append(entities, entity.PaymentEntity{
			ID:               val.ID,
			OrderID:          val.OrderID,
			UserID:           val.UserID,
			PaymentMethod:    val.PaymentMethod,
			PaymentStatus:    val.PaymentStatus,
			PaymentGatewayID: *val.PaymentGatewayID,
			GrossAmount:      val.GrossAmount,
			PaymentURL:       *val.PaymentURL,
		})
	}

	return entities, countData, int64(totalPage), nil
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
