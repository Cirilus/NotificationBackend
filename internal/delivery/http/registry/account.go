package registry

import (
	"github.com/gin-gonic/gin"
	"notification/internal/delivery/http/handlers"
)

func AccountRegistry(router *gin.RouterGroup, h *handlers.AccountHandler) {
	router.GET("/", h.GetAllAccounts)
	// router.GET("/me", h.GetMe)
	router.GET("/:id", h.GetAccountById)

	router.POST("/", h.CreateAccount)

	router.DELETE("/:id", h.DeleteAccount)

	router.PUT("/:id", h.UpdateAccount)
}
