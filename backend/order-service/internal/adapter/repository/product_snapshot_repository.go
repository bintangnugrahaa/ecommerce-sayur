package repository

import (
	"order-service/internal/core/domain/entity"
	"order-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProductSnapshotRepositoryInterface interface {
	Create(req entity.ProductConsumerResponse) error
	GetByID(productID int64) (*model.ProductSnapshot, error)
}

type ProductSnapshotRepository struct {
	db *gorm.DB
}

// Create implements ProductSnapshotRepositoryInterface.
func (p *ProductSnapshotRepository) Create(req entity.ProductConsumerResponse) error {
	modelProductSnapshot := model.ProductSnapshot{
		ID:           int64(req.ID),
		Name:         req.Name,
		Stock:        req.Stock,
		Image:        req.Image,
		RegulerPrice: req.RegularPrice,
		SalePrice:    req.SalePrice,
		Unit:         req.Unit,
		Weight:       req.Weight,
		CreatedAt:    req.CreatedAt,
	}

	if err := p.db.FirstOrCreate(&modelProductSnapshot, &model.ProductSnapshot{ID: int64(req.ID)}).Error; err != nil {
		log.Errorf("[ProductSnapshotRepository-1] Create: %v", err)
		return err
	}

	if len(req.Child) > 0 {
		for _, child := range req.Child {
			modelProductSnapshotChild := model.ProductSnapshot{
				ID:           int64(child.ID),
				Name:         child.Name,
				Stock:        child.Stock,
				Image:        child.Image,
				RegulerPrice: child.RegularPrice,
				SalePrice:    child.SalePrice,
				Unit:         child.Unit,
				Weight:       child.Weight,
				CreatedAt:    child.CreatedAt,
			}

			if err := p.db.FirstOrCreate(&modelProductSnapshotChild, &model.ProductSnapshot{ID: int64(child.ID)}).Error; err != nil {
				log.Errorf("[ProductSnapshotRepository-2] Create: %v", err)
				return err
			}
		}
	}

	return nil
}

// GetByID implements ProductSnapshotRepositoryInterface.
func (p *ProductSnapshotRepository) GetByID(productID int64) (*model.ProductSnapshot, error) {
	var productSnapshot model.ProductSnapshot

	if err := p.db.Where("id = ?", productID).First(&productSnapshot).Error; err != nil {
		log.Errorf("[ProductSnapshotRepository-1] GetByID: %v", err)
		return nil, err
	}

	return &productSnapshot, nil
}

func NewProductSnapshotRepository(db *gorm.DB) ProductSnapshotRepositoryInterface {
	return &ProductSnapshotRepository{
		db: db,
	}
}
