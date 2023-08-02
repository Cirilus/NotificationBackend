package middlewares

import (
	"go.uber.org/fx"
	"notification/internal/config"
	"notification/internal/middlewares/keycloak"
)

type Middlewares struct {
	KeycloakMiddleware *keycloak.Middleware
}

func NewMiddlewares(cfg *config.AppConfig) *Middlewares {

	return &Middlewares{
		KeycloakMiddleware: keycloak.NewMiddleware(cfg),
	}
}

var Module = fx.Module(
	"Middlewares",
	fx.Provide(NewMiddlewares),
)
