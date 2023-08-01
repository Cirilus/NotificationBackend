package postgres

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"notification/internal/models"
	"notification/pkg/clients/postgresql"
)

type NotificationRepositoryInterface interface {
	CreateNotification(ctx context.Context, notification models.Notification) error

	DeleteNotification(ctx context.Context, uuid uuid.UUID) error

	UpdateNotifications(ctx context.Context, uuid uuid.UUID, notification models.UpdatedNotification) (*models.UpdatedNotification, error)

	GetNotificationById(ctx context.Context, uuid uuid.UUID) (*models.Notification, error)
	GetAllNotifications(ctx context.Context) ([]models.Notification, error)
	GetNotificationsByUser(ctx context.Context, uuid uuid.UUID) ([]models.Notification, error)
}

type NotificationRepository struct {
	*postgresql.Postgres
	logger *logrus.Logger
}

func NewNotificationRepository(postgres *postgresql.Postgres, logger *logrus.Logger) *NotificationRepository {
	return &NotificationRepository{Postgres: postgres, logger: logger}
}

func (n NotificationRepository) CreateNotification(ctx context.Context, notification models.Notification) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationRepository) DeleteNotification(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationRepository) UpdateNotifications(ctx context.Context, uuid uuid.UUID, notification models.UpdatedNotification) (*models.UpdatedNotification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationRepository) GetNotificationById(ctx context.Context, uuid uuid.UUID) (*models.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationRepository) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationRepository) GetNotificationsByUser(ctx context.Context, uuid uuid.UUID) ([]models.Notification, error) {
	//TODO implement me
	panic("implement me")
}
