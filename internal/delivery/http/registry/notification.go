package registry

import (
	"github.com/gin-gonic/gin"
	"notification/internal/delivery/http/handlers"
)

func NotificationRegistry(router *gin.RouterGroup, h *handlers.NotificationHandler) {

	//router.POST("/", h.CreateNotification)
	//router.GET("/", h.GetAllNotifications)
	//router.GET("/:id", h.GetNotificationsById)
	//router.PUT("/:id", h.UpdateNotification)
	//router.DELETE("/:id", h.DeleteNotification)
	//
	//router.GET("/user/:id", h.GetNotificationsByUser)
	//router.POST("/user/:id", h.AddNotificationOnAccount)
	//router.DELETE("/user/:id", h.DeleteNotificationFromAccount)
}
