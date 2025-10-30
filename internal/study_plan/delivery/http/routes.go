package http

import (
	"github.com/labstack/echo/v4"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
)

func MapStudyPlanRoutes(g *echo.Group, h studyplan.Handler) {
	g.GET("", h.GetAll())
	g.GET("/:id", h.Get())

}
