package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mustaphalimar/prepilotapp-backend/internal/health"
)

func MapHealthRoutes(r *echo.Echo, h health.Handler) {
	r.GET("/status", h.CheckHealth)
	r.GET("/debug/cors", h.CheckCORS)
}
