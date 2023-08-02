package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	_ "notification/docs"
	"notification/internal/delivery/http/handlers"
	"notification/internal/middlewares"
)

func Endpoints(
	logger *logrus.Logger,
	router *gin.Engine,
	handlers *handlers.Handlers,
	md *middlewares.Middlewares,
) {
	logger.Info("Registry the endpoints")
	apiRoutes := router.Group("/api")
	{
		accountRoutes := apiRoutes.Group("/account")
		{
			AccountRegistry(accountRoutes, handlers.AccountHandler, md.KeycloakMiddleware)
		}

		notificationRoutes := apiRoutes.Group("/notification")
		{
			NotificationRegistry(notificationRoutes, handlers.NotificationHandler, md)
		}
	}

	router.GET("/status", handlers.HealthCheckHandler.Check)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

var Module = fx.Module(
	"Registry",
	fx.Invoke(Endpoints),
)
