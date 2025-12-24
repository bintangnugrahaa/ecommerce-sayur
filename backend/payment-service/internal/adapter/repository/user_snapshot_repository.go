package repository

import (
	"payment-service/internal/core/domain/entity"
	"payment-service/internal/core/domain/model"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserSnapshotRepositoryInterface interface {
	Create(userDetail *entity.ProfileHttpResponse) error
	GetByUserID(userID int64) (*model.UserSnapshot, error)
	UpdateLastUsed(userID int64) error
}

type UserSnapshotRepository struct {
	db *gorm.DB
}

// Create implements [UserSnapshotRepositoryInterface].
func (u *UserSnapshotRepository) Create(userDetail *entity.ProfileHttpResponse) error {
	modelUserSnapshot := model.UserSnapshot{
		UserID:   userDetail.ID,
		Name:     userDetail.Name,
		Email:    userDetail.Email,
		Phone:    userDetail.Phone,
		Address:  userDetail.Address,
		Photo:    userDetail.Photo,
		RoleName: userDetail.RoleName,
		Lat:      userDetail.Lat,
		Lng:      userDetail.Lng,
	}

	if err := u.db.FirstOrCreate(&modelUserSnapshot, &model.UserSnapshot{UserID: userDetail.ID}).Error; err != nil {
		log.Errorf("[UserSnapshotRepository-1] Create: %v", err)
		return err
	}

	return nil
}

// GetByUserID implements [UserSnapshotRepositoryInterface].
func (u *UserSnapshotRepository) GetByUserID(userID int64) (*model.UserSnapshot, error) {
	var userSnapshot model.UserSnapshot

	if err := u.db.Where("user_id = ?", userID).First(&userSnapshot).Error; err != nil {
		log.Errorf("[UserSnapshotRepository-1] GetByUserID: %v", err)
		return nil, err
	}

	return &userSnapshot, nil
}

// UpdateLastUsed implements [UserSnapshotRepositoryInterface].
func (u *UserSnapshotRepository) UpdateLastUsed(userID int64) error {
	now := time.Now()
	if err := u.db.Model(&model.UserSnapshot{}).Where("user_id = ?", userID).Update("last_used", now).Error; err != nil {
		log.Errorf("[UserSnapshotRepository-1] UpdateLastUsed: %v", err)
		return err
	}

	return nil
}

func NewUserSnapshotRepository(db *gorm.DB) UserSnapshotRepositoryInterface {
	return &UserSnapshotRepository{db: db}
}
