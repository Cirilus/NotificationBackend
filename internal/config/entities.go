package config

type Mode string

const (
	Dev  Mode = "dev"
	Prod Mode = "prod"
	Test Mode = "test"
)

const (
	DefaultAppConfigPath   = "configs/config.go.yaml"
	DefaultGooseConfigPath = "configs/config.goose.yaml"
)

type PostgreSQL struct {
	Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
	Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
	Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
	Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
	Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
}

type Keycloak struct {
	Url           string  `yaml:"url" env-required:"true"`
	Realm         string  `yaml:"realm" env-required:"true"`
	FullCertsPath *string `yaml:"fullCertsPath"`
}
