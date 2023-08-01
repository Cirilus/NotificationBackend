package repository

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"notification/internal/repository/postgres"
)

type Repositories struct {
	AccountRepository postgres.AccountRepository
}

func GetRepositories(logger *logrus.Logger) *Repositories {
	repo := new(Repositories)
	logger.Info("Creating the repository")
	return repo
}

var Module = fx.Module(
	"Repository",
	fx.Provide(GetRepositories),
)
