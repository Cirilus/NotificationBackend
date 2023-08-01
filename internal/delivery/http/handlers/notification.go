package handlers

import (
	"github.com/sirupsen/logrus"
	"notification/internal/usecase"
)

type NotificationHandler struct {
	uc     *usecase.NotificationUsecase
	logger *logrus.Logger
}

func NewNotificationHandler(uc *usecase.NotificationUsecase, logger *logrus.Logger) *NotificationHandler {
	return &NotificationHandler{uc: uc, logger: logger}
}
