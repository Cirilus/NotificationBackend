package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
	"sync"
)

type AppConfig struct {
	Mode       Mode       `yaml:"mode"`
	Host       string     `yaml:"host"`
	Port       string     `yaml:"port"`
	Postgresql PostgreSQL `yaml:"postgresql"`
	Keycloak   Keycloak   `yaml:"keycloak"`
	LogLevel   string     `yaml:"loglevel"`
}

var appConfig *AppConfig
var once sync.Once

func GetAppConfig() *AppConfig {
	once.Do(func() {
		pathToConfig := os.Getenv("path_to_app_config")
		logrus.Info("read application config")
		appConfig = &AppConfig{}
		if pathToConfig == "" {
			pathToConfig = DefaultAppConfigPath
		}
		if err := cleanenv.ReadConfig(pathToConfig, appConfig); err != nil {
			help, _ := cleanenv.GetDescription(appConfig, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})
	return appConfig
}

var Module = fx.Module(
	"config",
	fx.Provide(
		GetAppConfig,
	),
)
