package repository

import (
	"payment-service/internal/core/domain/entity"
	"payment-service/internal/core/domain/model"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderSnapshotRepositoryInterface interface {
	Create(orderDetail *entity.OrderDetailHttpResponse) error
	GetByOrderID(orderID int64) (*model.OrderSnapshot, error)
	UpdateLastUsed(orderID int64) error
}

type OrderSnapshotRepository struct {
	db *gorm.DB
}

// Create implements [OrderSnapshotRepositoryInterface].
func (o *OrderSnapshotRepository) Create(orderDetail *entity.OrderDetailHttpResponse) error {
	modelOrderSnapshot := model.OrderSnapshot{
		OrderID:         orderDetail.ID,
		OrderCode:       orderDetail.OrderCode,
		OrderDatetime:   orderDetail.OrderDatetime,
		Status:          orderDetail.Status,
		PaymentMethod:   orderDetail.PaymentMethod,
		ShippingFee:     orderDetail.ShippingFee,
		ShippingType:    orderDetail.ShippingType,
		Remarks:         orderDetail.Remarks,
		TotalAmount:     orderDetail.TotalAmount,
		CustomerName:    orderDetail.Customer.CustomerName,
		CustomerPhone:   orderDetail.Customer.CustomerPhone,
		CustomerAddress: orderDetail.Customer.CustomerAddress,
		CustomerEmail:   orderDetail.Customer.CustomerEmail,
		CustomerID:      orderDetail.Customer.CustomerID,
	}

	if err := o.db.FirstOrCreate(&modelOrderSnapshot, &model.OrderSnapshot{OrderID: orderDetail.ID}).Error; err != nil {
		log.Errorf("[OrderSnapshotRepository-1] Create: %v", err)
		return err
	}

	return nil
}

// GetByOrderID implements [OrderSnapshotRepositoryInterface].
func (o *OrderSnapshotRepository) GetByOrderID(orderID int64) (*model.OrderSnapshot, error) {
	var orderSnapshot model.OrderSnapshot

	if err := o.db.Where("order_id = ?", orderID).First(&orderSnapshot).Error; err != nil {
		log.Errorf("[OrderSnapshotRepository-1] GetByOrderId: %v", err)
		return nil, err
	}

	return &orderSnapshot, nil
}

// UpdateLastUsed implements [OrderSnapshotRepositoryInterface].
func (o *OrderSnapshotRepository) UpdateLastUsed(orderID int64) error {
	now := time.Now()
	if err := o.db.Model(&model.OrderSnapshot{}).Where("order_id = ?", orderID).Update("last_used", now).Error; err != nil {
		log.Errorf("[OrderSnapshotRepository-1] UpdateLastUsed: %v", err)
		return err
	}
	return nil
}

func NewOrderSnapshotRepository(db *gorm.DB) OrderSnapshotRepositoryInterface {
	return &OrderSnapshotRepository{db: db}
}
