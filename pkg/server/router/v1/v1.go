package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/handlers"
)

func RegisterV1Routes(router *echo.Group, handlers *handlers.Handlers) {

	registerStudyPlanRoutes(router, handlers.StudyPlan)
}
