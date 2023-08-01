package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
	"notification/internal/config"
)

func SetUpRouter(cfg *config.AppConfig) *gin.Engine {
	if cfg.Mode == config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.Use(gin.Recovery())

	return router
}

// @title Notification Service
// @version 1.0
// @description This is the documentation of the backend service of the notificator
// @host localhost:8000
// @BasePath /
// @schemes http
func RunApp(lc fx.Lifecycle, logger *logrus.Logger, router *gin.Engine, cfg *config.AppConfig) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					err := router.Run(cfg.Host + ":" + cfg.Port)
					if err != nil {
						logger.Errorf("There is an error in starting app")
						panic(err)
					}
				}()
				logger.Infof("Starting the app on the %s:%s", cfg.Host, cfg.Port)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info("Stopping the application")
				return nil
			},
		},
	)
}

var Module = fx.Module(
	"StartServer",
	fx.Provide(SetUpRouter),
	fx.Invoke(RunApp),
)
