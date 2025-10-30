package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mustaphalimar/prepilotapp-backend/internal/health"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
)

type HealthHandler struct {
	server *server.Server
}

func NewHealthHandler(server *server.Server) health.Handler {
	return &HealthHandler{
		server: server,
	}
}

func (h *HealthHandler) CheckCORS(c echo.Context) error {
	response := map[string]interface{}{
		"timestamp":   time.Now().UTC(),
		"environment": h.server.Config.Primary.Env,
		"cors_config": map[string]interface{}{
			"allowed_origins": h.server.Config.Server.CORSAllowedOrigins,
		},
		"request_headers": map[string]interface{}{
			"origin":     c.Request().Header.Get("Origin"),
			"user_agent": c.Request().Header.Get("User-Agent"),
			"referer":    c.Request().Header.Get("Referer"),
		},
		"note": "This endpoint helps debug CORS issues. Check if your frontend origin is in allowed_origins.",
	}

	return c.JSON(http.StatusOK, response)

}

func (h *HealthHandler) CheckHealth(c echo.Context) error {

	response := map[string]interface{}{
		"status":      "healthy",
		"timestamp":   time.Now().UTC(),
		"environment": h.server.Config.Primary.Env,
		"checks":      make(map[string]interface{}),
	}

	checks := response["checks"].(map[string]interface{})
	isHealthy := true

	// Check database connectivity
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbStart := time.Now()
	if err := h.server.DB.PgxPool.Ping(ctx); err != nil {
		checks["database"] = map[string]interface{}{
			"status":        "unhealthy",
			"response_time": time.Since(dbStart).String(),
			"error":         err.Error(),
		}
		isHealthy = false
		log.Fatalln("database health check failed")
	} else {
		checks["database"] = map[string]interface{}{
			"status":        "healthy",
			"response_time": time.Since(dbStart).String(),
		}
		log.Println("database health check passed")
	}

	// Database connection metrics are automatically captured by New Relic nrpgx5 integration

	// Check Redis connectivity
	if h.server.Redis != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		redisStart := time.Now()
		if err := h.server.Redis.Ping(ctx).Err(); err != nil {
			checks["redis"] = map[string]interface{}{
				"status":        "unhealthy",
				"response_time": time.Since(redisStart).String(),
				"error":         err.Error(),
			}
			log.Fatalln("redis health check failed")
		} else {
			checks["redis"] = map[string]interface{}{
				"status":        "healthy",
				"response_time": time.Since(redisStart).String(),
			}
			log.Println("redis health check passed")
		}
	}

	// Set overall status
	if !isHealthy {
		response["status"] = "unhealthy"
		log.Println("health check failed")
		return c.JSON(http.StatusServiceUnavailable, response)
	}

	log.Println("health check passed")

	err := c.JSON(http.StatusOK, response)
	if err != nil {
		log.Fatalln("failed to write JSON response")

		return fmt.Errorf("failed to write JSON response: %w", err)
	}

	return nil
}
