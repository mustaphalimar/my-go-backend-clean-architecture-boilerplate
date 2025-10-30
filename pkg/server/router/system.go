package router

import (
	"github.com/mustaphalimar/prepilotapp-backend/internal/health/delivery/http"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/handlers"

	"github.com/labstack/echo/v4"
)

func registerSystemRoutes(r *echo.Echo, h *handlers.Handlers) {
	http.MapHealthRoutes(r, h.Health)
	// r.Static("/static", "static")

}
