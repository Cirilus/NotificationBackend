package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
	"sync"
)

type Config struct {
	Mode       Mode       `yaml:"mode"`
	Host       string     `yaml:"host"`
	Port       string     `yaml:"port"`
	Postgresql PostgreSQL `yaml:"postgresql"`
	Keycloak   Keycloak   `yaml:"keycloak"`
	LogLevel   string     `yaml:"loglevel"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		pathToConfig := os.Getenv("path_to_config")
		logrus.Info("read application config")
		instance = &Config{}
		if pathToConfig == "" {
			pathToConfig = DefaultConfigPath
		}
		if err := cleanenv.ReadConfig(pathToConfig, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})
	return instance
}

var Module = fx.Module(
	"config",
	fx.Provide(
		GetConfig,
	),
)
