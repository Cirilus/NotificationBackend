package middlewares

import (
	"go.uber.org/fx"
	"notification/internal/middlewares/keycloak"
)

type Middlewares struct {
	KeycloakMiddleware *keycloak.Middleware
}

func NewMiddlewares() *Middlewares {
	keycloakMiddleware := keycloak.Middleware{}

	return &Middlewares{
		KeycloakMiddleware: &keycloakMiddleware,
	}
}

var Module = fx.Module(
	"Middlewares",
	fx.Provide(NewMiddlewares),
)
