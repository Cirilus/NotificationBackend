package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"notification/internal/config"
)

func SetUpRouter(logger *logrus.Logger, cfg *config.Config) *gin.Engine {
	if cfg.Mode == config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	return router
}

func RunApp(lc fx.Lifecycle, logger *logrus.Logger, router *gin.Engine, cfg *config.Config) {
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
