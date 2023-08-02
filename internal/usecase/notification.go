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

	DeleteNotification(ctx context.Context, id uuid.UUID) error

	UpdateNotifications(ctx context.Context, id uuid.UUID, notification models.UpdatedNotification) (*models.UpdatedNotification, error)

	NotificationById(ctx context.Context, id uuid.UUID) (*models.Notification, error)
	AllNotifications(ctx context.Context) ([]models.Notification, error)
	NotificationsByAccount(ctx context.Context, id uuid.UUID) ([]models.Notification, error)
}

type NotificationUsecase struct {
	repo   postgres.NotificationRepositoryInterface
	logger *logrus.Logger
}

func NewNotificationUsecase(repo postgres.NotificationRepositoryInterface, logger *logrus.Logger) *NotificationUsecase {
	return &NotificationUsecase{repo: repo, logger: logger}
}

func (n NotificationUsecase) CreateNotification(ctx context.Context, notification models.Notification) error {
	err := n.repo.CreateNotification(ctx, notification)
	if err != nil {
		logrus.Error("Notification - UseCase - CreateNotification")
		return err
	}
	return nil
}

func (n NotificationUsecase) AddNotificationToAccount(ctx context.Context, notificationId uuid.UUID, accountsId []uuid.UUID) error {
	//TODO implement me from creating the Add Account to Notification in repo
	panic("implement me")
}

func (n NotificationUsecase) DeleteNotificationFromAccount(ctx context.Context, notificationId uuid.UUID, accountId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUsecase) DeleteNotification(ctx context.Context, id uuid.UUID) error {
	err := n.repo.DeleteNotification(ctx, id)
	if err != nil {
		logrus.Error("Notification - UseCase - DeleteNotification")
		return err
	}
	return nil
}

func (n NotificationUsecase) UpdateNotifications(ctx context.Context, id uuid.UUID, notification models.UpdatedNotification) (*models.UpdatedNotification, error) {
	updateNotifications, err := n.repo.UpdateNotifications(ctx, id, notification)
	if err != nil {
		logrus.Error("Notification - UseCase - UpdateNotification")
		return nil, err
	}
	return updateNotifications, nil
}

func (n NotificationUsecase) NotificationById(ctx context.Context, id uuid.UUID) (*models.Notification, error) {
	notificationById, err := n.repo.GetNotificationById(ctx, id)
	if err != nil {
		logrus.Error("Notification - UseCase - NotificationById")
		return nil, err
	}
	return notificationById, nil
}

func (n NotificationUsecase) AllNotifications(ctx context.Context) ([]models.Notification, error) {
	allNotifications, err := n.repo.GetAllNotifications(ctx)
	if err != nil {
		logrus.Error("Notification - UseCase - AllNotifications")
		return nil, err
	}
	return allNotifications, nil
}

func (n NotificationUsecase) NotificationsByAccount(ctx context.Context, id uuid.UUID) ([]models.Notification, error) {
	notificationByUser, err := n.repo.GetNotificationsByAccount(ctx, id)
	if err != nil {
		logrus.Error("Notification - UseCase - NotificationById")
		return nil, err
	}
	return notificationByUser, nil
}
