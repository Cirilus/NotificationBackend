package registry

import (
	"github.com/gin-gonic/gin"
	"notification/internal/delivery/http/handlers"
	"notification/internal/middlewares"
)

func NotificationRegistry(
	router *gin.RouterGroup,
	h *handlers.NotificationHandler,
	m *middlewares.Middlewares,
) {

	router.GET("/", h.AllNotifications)
	router.GET("/:id", h.NotificationById)
	router.GET("/user/:id", h.NotificationsByAccount)

	router.PUT("/:id", h.UpdateNotification)

	router.DELETE("/:id", h.DeleteNotification)
	router.DELETE("/user/:id", h.DeleteNotificationFromAccount)

	router.POST("/", h.CreateNotification)
	router.POST("/user/:id", h.AddNotificationToAccount)

}
