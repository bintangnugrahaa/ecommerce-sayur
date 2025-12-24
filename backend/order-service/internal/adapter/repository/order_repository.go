package repository

import (
	"context"
	"errors"
	"math"
	"order-service/internal/core/domain/entity"
	"order-service/internal/core/domain/model"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	GetAll(ctx context.Context, queryString entity.QueryStringEntity) ([]entity.OrderEntity, int64, int64, error)
	GetByID(ctx context.Context, orderID int64) (*entity.OrderEntity, error)
	CreateOrder(ctx context.Context, req entity.OrderEntity) (int64, error)
	UpdateStatus(ctx context.Context, req entity.OrderEntity) (int64, string, string, error)
	DeleteOrder(ctx context.Context, orderID int64) error

	GetOrderByOrderCode(ctx context.Context, orderCode string) (*entity.OrderEntity, error)
	GetProductFromSnapshot(productID int64) (*entity.ProductResponseEntity, error)
}

type orderRepository struct {
	db                  *gorm.DB
	productSnapshotRepo ProductSnapshotRepositoryInterface
}

// GetProductFromSnapshot implements [OrderRepositoryInterface].
func (o *orderRepository) GetProductFromSnapshot(productID int64) (*entity.ProductResponseEntity, error) {
	productSnapshot, err := o.productSnapshotRepo.GetByID(productID)
	if err != nil {
		log.Errorf("[OrderRepository-1] GetProductFromSnapshot: %v", err)
		return nil, err
	}

	return &entity.ProductResponseEntity{
		ID:           int(productSnapshot.ID),
		Stock:        productSnapshot.Stock,
		ProductName:  productSnapshot.Name,
		ProductImage: productSnapshot.Image,
		SalePrice:    float64(productSnapshot.SalePrice),
		Unit:         productSnapshot.Unit,
		Weight:       productSnapshot.Weight,
	}, nil
}

// GetOrderByOrderCode implements OrderRepositoryInterface.
func (o *orderRepository) GetOrderByOrderCode(ctx context.Context, orderCode string) (*entity.OrderEntity, error) {
	var modelOrder model.Order

	if err := o.db.Preload("OrderItems").Where("order_code =?", orderCode).First(&modelOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Infof("[OrderRepository-1] GetOrderByOrderCode: Order not found")
			return nil, err
		}
		log.Errorf("[OrderRepository-2] GetOrderByOrderCode: %v", err)
		return nil, err
	}

	orderItemEntities := []entity.OrderItemEntity{}
	for _, item := range modelOrder.OrderItems {
		orderItemEntities = append(orderItemEntities, entity.OrderItemEntity{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return &entity.OrderEntity{
		ID:           modelOrder.ID,
		OrderCode:    modelOrder.OrderCode,
		Status:       modelOrder.Status,
		BuyerID:      modelOrder.BuyerID,
		OrderDate:    modelOrder.OrderDate.Format("2006-01-02 15:04:05"),
		TotalAmount:  int64(modelOrder.TotalAmount),
		OrderItems:   orderItemEntities,
		Remarks:      modelOrder.Remarks,
		ShippingType: modelOrder.ShippingType,
		ShippingFee:  int64(modelOrder.ShippingFee),
		BuyerEmail:   modelOrder.BuyerEmail,
		BuyerPhone:   modelOrder.BuyerPhone,
		BuyerAddress: modelOrder.BuyerAddress,
		BuyerLat:     modelOrder.BuyerLat,
		BuyerLng:     modelOrder.BuyerLng,
	}, nil
}

// CreateOrder implements OrderRepositoryInterface.
func (o *orderRepository) CreateOrder(ctx context.Context, req entity.OrderEntity) (int64, error) {
	orderDate, err := time.Parse("2006-01-02", req.OrderDate) // YYYY-MM-DD
	if err != nil {
		log.Errorf("[OrderRepository-1] CreateOrder: %v", err)
		return 0, err
	}

	var orderItems []model.OrderItem
	for _, item := range req.OrderItems {
		orderItem := model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}

	modelOrder := model.Order{
		OrderCode:    req.OrderCode,
		BuyerID:      req.BuyerID,
		OrderDate:    orderDate,
		OrderTime:    req.OrderTime,
		Status:       req.Status,
		TotalAmount:  float64(req.TotalAmount),
		ShippingType: req.ShippingType,
		ShippingFee:  float64(req.ShippingFee),
		Remarks:      req.Remarks,
		OrderItems:   orderItems,
		BuyerName:    req.BuyerName,
		BuyerEmail:   req.BuyerEmail,
		BuyerPhone:   req.BuyerPhone,
		BuyerAddress: req.BuyerAddress,
		BuyerLat:     req.BuyerLat,
		BuyerLng:     req.BuyerLng,
	}

	if err := o.db.Create(&modelOrder).Error; err != nil {
		log.Errorf("[OrderRepository-3] CreateOrder: %v", err)
		return 0, err
	}

	return modelOrder.ID, nil
}

// DeleteOrder implements OrderRepositoryInterface.
func (o *orderRepository) DeleteOrder(ctx context.Context, orderID int64) error {
	modelOrder := model.Order{}

	if err := o.db.Preload("OrderItems").Where("id = ?", orderID).First(&modelOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Infof("[OrderRepository-1] DeleteOrder: Order not found")
			return err
		}
		log.Errorf("[OrderRepository-2] DeleteOrder: %v", err)
		return err
	}

	if err := o.db.Select("OrderItems").Delete(&modelOrder).Error; err != nil {
		log.Errorf("[OrderRepository-3] DeleteOrder: %v", err)
		return err
	}

	return nil
}

// EditOrder implements OrderRepositoryInterface.
func (o *orderRepository) UpdateStatus(ctx context.Context, req entity.OrderEntity) (int64, string, string, error) {
	modelOrder := model.Order{}

	if err := o.db.Select("id", "order_code", "status", "buyer_id", "remarks").Where("id = ?", req.ID).First(&modelOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Infof("[OrderRepository-1] UpdateStatus: Order not found")
			return 0, "", "", err
		}
		log.Errorf("[OrderRepository-2] UpdateStatus: %v", err)
		return 0, "", "", err
	}

	if modelOrder.Status == "Pending" && (req.Status != "Confirmed" && req.Status != "Cancelled") {
		log.Infof("[OrderRepository-3] UpdateStatus: Invalid status transition")
		return 0, "", "", errors.New("400")
	}

	if modelOrder.Status == "Confirmed" && (req.Status != "Process" && req.Status != "Cancelled") {
		log.Infof("[OrderRepository-4] UpdateStatus: Invalid status transition")
		return 0, "", "", errors.New("400")
	}

	if modelOrder.Status == "Process" && (req.Status != "Sending" && req.Status != "Cancelled") {
		log.Infof("[OrderRepository-5] UpdateStatus: Invalid status transition")
		return 0, "", "", errors.New("400")
	}

	if modelOrder.Status == "Sending" && (req.Status != "Done" && req.Status != "Cancelled") {
		log.Infof("[OrderRepository-6] UpdateStatus: Invalid status transition")
		return 0, "", "", errors.New("400")
	}

	modelOrder.Status = req.Status
	modelOrder.Remarks = req.Remarks

	if err := o.db.UpdateColumns(&modelOrder).Error; err != nil {
		log.Errorf("[OrderRepository-7] UpdateStatus: %v", err)
		return 0, "", "", err
	}

	return modelOrder.BuyerID, modelOrder.Status, modelOrder.OrderCode, nil
}

// GetAll implements OrderRepositoryInterface.
func (o *orderRepository) GetAll(ctx context.Context, queryString entity.QueryStringEntity) ([]entity.OrderEntity, int64, int64, error) {
	var modelOrders []model.Order
	var countData int64
	offset := (queryString.Page - 1) * queryString.Limit

	sqlMain := o.db.Preload("OrderItems").
		Where("order_code ILIKE ? OR status ILIKE ?", "%"+queryString.Search+"%", "%"+queryString.Status+"%")

	if queryString.BuyerID != 0 {
		sqlMain = sqlMain.Where("buyer_id = ?", queryString.BuyerID)
	}

	if err := sqlMain.Model(&modelOrders).Count(&countData).Error; err != nil {
		log.Errorf("[OrderRepository-1] GetAll: %v", err)
		return nil, 0, 0, err
	}

	totalPage := int(math.Ceil(float64(countData) / float64(queryString.Limit)))
	if err := sqlMain.Order("order_date DESC").Limit(int(queryString.Limit)).Offset(int(offset)).Find(&modelOrders).Error; err != nil {
		log.Errorf("[OrderRepository-2] GetAll: %v", err)
		return nil, 0, 0, err
	}

	if len(modelOrders) == 0 {
		err := errors.New("404")
		log.Infof("[OrderRepository-3] GetAll: No order found")
		return nil, 0, 0, err
	}

	entities := []entity.OrderEntity{}
	for _, val := range modelOrders {
		orderItemEntities := []entity.OrderItemEntity{}
		for _, item := range val.OrderItems {
			orderItemEntities = append(orderItemEntities, entity.OrderItemEntity{
				ID:        item.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
		}
		entities = append(entities, entity.OrderEntity{
			ID:           val.ID,
			OrderCode:    val.OrderCode,
			Status:       val.Status,
			OrderDate:    val.OrderDate.Format("2006-01-02 15:04:05"),
			TotalAmount:  int64(val.TotalAmount),
			OrderItems:   orderItemEntities,
			BuyerID:      val.BuyerID,
			BuyerName:    val.BuyerName,
			BuyerEmail:   val.BuyerEmail,
			BuyerPhone:   val.BuyerPhone,
			BuyerAddress: val.BuyerAddress,
			BuyerLat:     val.BuyerLat,
			BuyerLng:     val.BuyerLng,
		})
	}

	return entities, countData, int64(totalPage), nil
}

// GetByID implements OrderRepositoryInterface.
func (o *orderRepository) GetByID(ctx context.Context, orderID int64) (*entity.OrderEntity, error) {
	var modelOrder model.Order

	if err := o.db.Preload("OrderItems").Where("id =?", orderID).First(&modelOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Infof("[OrderRepository-1] GetByID: Order not found")
			return nil, err
		}
		log.Errorf("[OrderRepository-2] GetByID: %v", err)
		return nil, err
	}

	orderItemEntities := []entity.OrderItemEntity{}
	for _, item := range modelOrder.OrderItems {
		orderItemEntities = append(orderItemEntities, entity.OrderItemEntity{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return &entity.OrderEntity{
		ID:           modelOrder.ID,
		OrderCode:    modelOrder.OrderCode,
		Status:       modelOrder.Status,
		BuyerID:      modelOrder.BuyerID,
		OrderDate:    modelOrder.OrderDate.Format("2006-01-02 15:04:05"),
		TotalAmount:  int64(modelOrder.TotalAmount),
		OrderItems:   orderItemEntities,
		Remarks:      modelOrder.Remarks,
		ShippingType: modelOrder.ShippingType,
		ShippingFee:  int64(modelOrder.ShippingFee),
		BuyerName:    modelOrder.BuyerName,
		BuyerEmail:   modelOrder.BuyerEmail,
		BuyerPhone:   modelOrder.BuyerPhone,
		BuyerAddress: modelOrder.BuyerAddress,
		BuyerLat:     modelOrder.BuyerLat,
		BuyerLng:     modelOrder.BuyerLng,
	}, nil
}

func NewOrderRepository(db *gorm.DB, productSnapshot ProductSnapshotRepositoryInterface) OrderRepositoryInterface {
	return &orderRepository{db: db, productSnapshotRepo: productSnapshot}
}
