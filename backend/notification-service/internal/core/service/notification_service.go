package service

import (
	"context"
	"notification-service/internal/adapter/repository"
	"notification-service/internal/core/domain/entity"
)

type NotificationServiceInterface interface {
	GetAll(ctx context.Context, queryString entity.NotifyQueryString) ([]entity.NotificationEntity, int64, int64, error)
	GetByID(ctx context.Context, notifID uint) (*entity.NotificationEntity, error)
}

type NotificationService struct {
	repo repository.NotificationRepositoryInterface
}

// GetByID implements [NotificationServiceInterface].
func (n *NotificationService) GetByID(ctx context.Context, notifID uint) (*entity.NotificationEntity, error) {
	return n.repo.GetByID(ctx, notifID)
}

// GetAll implements [NotificationServiceInterface].
func (n *NotificationService) GetAll(ctx context.Context, queryString entity.NotifyQueryString) ([]entity.NotificationEntity, int64, int64, error) {
	return n.repo.GetAll(ctx, queryString)
}

func NewNotificationService(repo repository.NotificationRepositoryInterface) NotificationServiceInterface {
	return &NotificationService{repo: repo}
}
