package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	_ "notification/docs"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

var Module = fx.Module(
	"Registry",
	fx.Invoke(Endpoints),
)
