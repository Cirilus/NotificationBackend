package repository

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"notification/internal/repository/postgres"
	client "notification/pkg/clients/postgresql"
)

type Repositories struct {
	AccountRepository      *postgres.AccountRepository
	NotificationRepository *postgres.NotificationRepository
}

func GetRepositories(logger *logrus.Logger, client *client.Postgres) *Repositories {
	logger.Info("Creating the repository")
	repo := Repositories{
		AccountRepository:      postgres.NewAccountRepository(client, logger),
		NotificationRepository: postgres.NewNotificationRepository(client, logger),
	}
	return &repo
}

var Module = fx.Module(
	"Repository",
	fx.Provide(GetRepositories),
)
