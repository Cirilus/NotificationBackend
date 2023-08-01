package usecase

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"notification/internal/repository"
)

type Usecase struct {
	AccountUsecase      *AccountUsecase
	NotificationUsecase *NotificationUsecase
}

func GetUsecases(logger *logrus.Logger, repo *repository.Repositories) *Usecase {
	logger.Infof("Creating the usecase")
	usecase := Usecase{
		AccountUsecase:      NewAccountUsecase(repo.AccountRepository, logger),
		NotificationUsecase: NewNotificationUsecase(repo.NotificationRepository, logger),
	}

	return &usecase
}

var Module = fx.Module(
	"Usecases",
	fx.Provide(GetUsecases),
)
