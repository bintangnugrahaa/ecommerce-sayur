package repository

import (
	"context"
	"errors"
	"fmt"
	"math"
	"notification-service/internal/core/domain/entity"
	"notification-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	GetAll(ctx context.Context, queryString entity.NotifyQueryString) ([]entity.NotificationEntity, int64, int64, error)
}

type notificationRepository struct {
	db *gorm.DB
}

// GetAll implements [NotificationRepositoryInterface].
func (n *notificationRepository) GetAll(ctx context.Context, queryString entity.NotifyQueryString) ([]entity.NotificationEntity, int64, int64, error) {
	modelNotifes := []model.Notification{}

	var countData int64
	offset := (queryString.Page - 1) * queryString.Limit

	sqlMain := n.db.
		Select("id", "subject", "status", "sent_at").
		Where("subject ILIKE ? OR message ILIKE ? OR status ILIKE ?", "%"+queryString.Search+"%", "%"+queryString.Search+"%", "%"+queryString.Status+"%")

	if queryString.UserID != 0 {
		sqlMain = sqlMain.Where("reciever_id = ?", queryString.UserID)
	}

	if queryString.IsRead {
		sqlMain = sqlMain.Where("read_at IS NOT NULL")
	}

	if err := sqlMain.Model(&modelNotifes).Count(&countData).Error; err != nil {
		log.Errorf("[NotificationRepository-1] GetAll: %v", err)
		return nil, 0, 0, err
	}

	order := fmt.Sprintf("%s %s", queryString.OrderBy, queryString.OrderType)

	totalPage := int(math.Ceil(float64(countData) / float64(queryString.Limit)))
	if err := sqlMain.Order(order).Limit(int(queryString.Limit)).Offset(int(offset)).Find(&modelNotifes).Error; err != nil {
		log.Errorf("[NotificationRepository-2] GetAll: %v", err)
		return nil, 0, 0, err
	}

	if len(modelNotifes) == 0 {
		err := errors.New("404")
		log.Infof("[NotificationRepository-3] GetAll: No notification found")
		return nil, 0, 0, err
	}
	notifEntities := []entity.NotificationEntity{}
	for _, val := range modelNotifes {
		notifEntities = append(notifEntities, entity.NotificationEntity{
			ID:      val.ID,
			Subject: val.Subject,
			Status:  val.Status,
			SentAt:  val.SentAt,
		})
	}

	return notifEntities, countData, int64(totalPage), nil
}

func NewNotificationRepository(db *gorm.DB) NotificationRepositoryInterface {
	return &notificationRepository{db: db}
}
