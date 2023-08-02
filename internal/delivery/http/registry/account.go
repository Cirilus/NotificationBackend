package registry

import (
	"github.com/gin-gonic/gin"
	"notification/internal/delivery/http/handlers"
	"notification/internal/middlewares/keycloak"
)

func AccountRegistry(
	router *gin.RouterGroup,
	h *handlers.AccountHandler,
	km *keycloak.Middleware,
) {

	protected := router.Use(km.Auth(km.AuthCheck(), km.Config))
	{
		protected.GET("/", h.GetAllAccounts)
		protected.GET("/:id", h.GetAccountById)
		protected.DELETE("/:id", h.DeleteAccount)
		protected.PUT("/:id", h.UpdateAccount)
		// protected.GET("/me", h.GetMe)
	}

	router.POST("/", h.CreateAccount)

}
