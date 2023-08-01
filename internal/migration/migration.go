package migration

import (
	"github.com/pressly/goose"
	"notification/internal/config"
	"notification/pkg/clients/postgresql"
)

func Migrate(postgres *postgresql.Postgres, cfg *config.GooseConfig) {
	err := goose.SetDialect(cfg.Dialect)
	if err != nil {
		panic(err)
	}

}
