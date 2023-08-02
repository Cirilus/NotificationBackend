package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"notification/internal/models"
	"notification/internal/usecase"
)

type NotificationHandler struct {
	uc     *usecase.NotificationUsecase
	logger *logrus.Logger
}

func NewNotificationHandler(uc *usecase.NotificationUsecase, logger *logrus.Logger) *NotificationHandler {
	return &NotificationHandler{uc: uc, logger: logger}
}

// CreateNotification swaggerDoc
// @Summary Create the User's notifications.
// @Tags notification
// @Accept  json
// @Produce json
// @Param notification body models.Notification true "Notification object that needs to be created"
// @Success 201
// @Failure 500
// @Router /api/notification/ [post]
func (h NotificationHandler) CreateNotification(c *gin.Context) {
	inp := new(models.Notification)
	err := c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request, err= %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.uc.CreateNotification(c.Request.Context(), *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

// AddNotificationToAccount swaggerDoc
// @Summary		Add notification to the account
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "notification id"
// @Param		id body []string true "account id"
// @Success 	201
// @Failure 	500
// @Router		/api/notification/user/{id} [post]
func (h NotificationHandler) AddNotificationToAccount(c *gin.Context) {
	notificationId := c.Param("id")
	notificationUUID, err := uuid.Parse(notificationId)

	if err != nil {
		logrus.Errorf("Cannot parse notification id to uuid, err= %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	inp := new([]uuid.UUID)
	err = c.BindJSON(inp)

	if err != nil {
		logrus.Errorf("Bad request, err = %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.uc.AddNotificationToAccount(c.Request.Context(), notificationUUID, *inp)

	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteNotificationFromAccount swaggerDoc
// @Summary		Delete the notification from account
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "notification id"
// @Param		id body string true "account id"
// @Success 	200
// @Failure 	500
// @Router		/api/notification/user/{id} [delete]
func (h NotificationHandler) DeleteNotificationFromAccount(c *gin.Context) {
	notificationId := c.Param("id")
	notificationUUID, err := uuid.Parse(notificationId)
	if err != nil {
		logrus.Errorf("Cannot parse notification id to uuid, err= %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	inp := new(uuid.UUID)
	err = c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request, err= %s", err)
		return
	}

	err = h.uc.DeleteNotificationFromAccount(c.Request.Context(), notificationUUID, *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// DeleteNotification swaggerDoc
// @Summary		Delete the notification
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "delete the notifications"
// @Success 	200
// @Failure 	500
// @Router		/api/notification/{id} [delete]
func (h NotificationHandler) DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = h.uc.DeleteNotification(c.Request.Context(), accountUUID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// UpdateNotification swaggerDoc
// @Summary		Update the notification
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "id of the notifications"
// @Param		updatednotification body models.Notification true "The Update notification"
// @Success 	200 {object} models.Notification
// @Failure 	500
// @Router		/api/notification/{id} [put]
func (h NotificationHandler) UpdateNotification(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	inp := new(models.UpdatedNotification)
	err = c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request, err= %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updateNotifications, err := h.uc.UpdateNotifications(c.Request.Context(), accountUUID, *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, updateNotifications)
}

// NotificationById swaggerDoc
// @Summary		Return notification by id
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "the id of the notification"
// @Success 	201 {object} models.Notification
// @Failure 	500
// @Router		/api/notification/{id} [get]
func (h NotificationHandler) NotificationById(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	notifications, err := h.uc.NotificationById(c.Request.Context(), accountUUID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notifications)
}

// AllNotifications swaggerDoc
// @Summary Return all notifications
// @Tags notification
// @Accept  json
// @Produce json
// @Success 201 {array} models.Notification
// @Failure 500
// @Router /api/notification/ [get]
func (h NotificationHandler) AllNotifications(c *gin.Context) {
	notifications, err := h.uc.AllNotifications(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notifications)
}

// NotificationsByAccount swaggerDoc
// @Summary		Return notification by account_id
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "user id of the notifications"
// @Success 	201 {array} models.Notification
// @Failure 	500
// @Router		/api/notification/user/{id} [get]
func (h NotificationHandler) NotificationsByAccount(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	notifications, err := h.uc.NotificationsByAccount(c.Request.Context(), accountUUID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notifications)
}
