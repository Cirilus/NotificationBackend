package handlers

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"notification/internal/usecase"
)

type Handlers struct {
	AccountHandler      *AccountHandler
	HealthCheckHandler  *HealthCheckHandler
	NotificationHandler *NotificationHandler
}

func GetHandlers(logger *logrus.Logger, uc *usecase.Usecase) *Handlers {
	logger.Info("Creating the handlers")
	handlers := Handlers{
		AccountHandler:      NewHandler(uc.AccountUsecase, logger),
		HealthCheckHandler:  NewHealthCheckHandler(),
		NotificationHandler: NewNotificationHandler(uc.NotificationUsecase, logger),
	}
	return &handlers
}

var Module = fx.Module(
	"Handlers",
	fx.Provide(GetHandlers),
)
