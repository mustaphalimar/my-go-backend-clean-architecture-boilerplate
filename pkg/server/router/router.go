package router

import (
	"github.com/labstack/echo/v4"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/handlers"
	v1 "github.com/mustaphalimar/prepilotapp-backend/pkg/server/router/v1"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/usecases"
)

func NewRouter(srv *server.Server, handlers *handlers.Handlers, usecases *usecases.Usecases) *echo.Echo {
	router := echo.New()

	// api versionning
	v1Router := router.Group("/api/v1")

	v1.RegisterV1Routes(v1Router, handlers)
	return router
}
