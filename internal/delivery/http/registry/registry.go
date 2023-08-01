package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"notification/internal/delivery/http/handlers"
)

func Endpoints(logger *logrus.Logger, router *gin.Engine, handlers *handlers.Handlers) {
	logger.Info("Registry the endpoints")
	apiRoutes := router.Group("/api")
	{
		accountRoutes := apiRoutes.Group("/account")
		{
			AccountRegistry(accountRoutes, handlers.AccountHandler)
		}
	}

	router.GET("/status", handlers.HealthCheckHandler.Check)
}

var Module = fx.Module(
	"Registry",
	fx.Invoke(Endpoints),
)
