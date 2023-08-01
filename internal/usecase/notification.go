package usecase

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"notification/internal/models"
	"notification/internal/repository/postgres"
)

type NotificationUseCaseInterface interface {
	CreateNotification(ctx context.Context, notification models.Notification) error

	AddNotificationToAccount(ctx context.Context, notificationId uuid.UUID, accountsId []uuid.UUID) error
	DeleteNotificationFromAccount(ctx context.Context, notificationId uuid.UUID, accountId uuid.UUID) error

	DeleteNotification(ctx context.Context, uuid uuid.UUID) error

	UpdateNotifications(ctx context.Context, uuid uuid.UUID, notification models.UpdatedNotification) (*models.UpdatedNotification, error)

	NotificationById(ctx context.Context, uuid uuid.UUID) (*models.Notification, error)
	AllNotifications(ctx context.Context) ([]models.Notification, error)
	NotificationsByUser(ctx context.Context, uuid uuid.UUID) ([]models.Notification, error)
}

type NotificationUsecase struct {
	repo   postgres.NotificationRepositoryInterface
	logger *logrus.Logger
}

func NewNotificationUsecase(repo postgres.NotificationRepositoryInterface, logger *logrus.Logger) *NotificationUsecase {
	return &NotificationUsecase{repo: repo, logger: logger}
}

func (n NotificationUsecase) CreateNotification(ctx context.Context, notification models.Notification) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) AddNotificationToAccount(ctx context.Context, notificationId uuid.UUID, accountsId []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) DeleteNotificationFromAccount(ctx context.Context, notificationId uuid.UUID, accountId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) DeleteNotification(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) UpdateNotifications(ctx context.Context, uuid uuid.UUID, notification models.UpdatedNotification) (*models.UpdatedNotification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) NotificationById(ctx context.Context, uuid uuid.UUID) (*models.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) AllNotifications(ctx context.Context) ([]models.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) NotificationsByUser(ctx context.Context, uuid uuid.UUID) ([]models.Notification, error) {
	//TODO implement me
	panic("implement me")
}
