package main

import (
	"go.uber.org/fx"
	"notification/internal/config"
	"notification/internal/delivery/http/handlers"
	"notification/internal/delivery/http/registry"
	"notification/internal/repository"
	"notification/internal/usecase"
	"notification/pkg/logger"
	"notification/server"
)

func main() {
	fx.New(
		config.Module,
		logger.Module,
		repository.Module,
		usecase.Module,
		handlers.Module,
		registry.Module,
		server.Module,
	).Run()
}
