package service

import (
	"context"
	"notification-service/internal/adapter/repository"
	"notification-service/internal/core/domain/entity"
)

type NotificationServiceInterface interface {
	GetAll(ctx context.Context, queryString entity.NotifyQueryString) ([]entity.NotificationEntity, int64, int64, error)
}

type NotificationService struct {
	repo repository.NotificationRepositoryInterface
}

// GetAll implements [NotificationServiceInterface].
func (n *NotificationService) GetAll(ctx context.Context, queryString entity.NotifyQueryString) ([]entity.NotificationEntity, int64, int64, error) {
	return n.repo.GetAll(ctx, queryString)
}

func NewNotificationService(repo repository.NotificationRepositoryInterface) NotificationServiceInterface {
	return &NotificationService{repo: repo}
}
