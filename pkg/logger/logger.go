package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"notification/internal/config"
	"os"
	"time"
)

func GetLogger(cfg *config.AppConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC822,
	})

	return logger
}

func GetLoggerFx(logger *logrus.Logger) fxevent.Logger {
	return &fxevent.ConsoleLogger{W: logger.Writer()}
}

var Module = fx.Module(
	"logger",
	fx.Provide(GetLogger),
	fx.WithLogger(GetLoggerFx),
)
