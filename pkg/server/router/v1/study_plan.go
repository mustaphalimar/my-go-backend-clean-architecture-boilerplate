package v1

import (
	"github.com/labstack/echo/v4"

	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
	"github.com/mustaphalimar/prepilotapp-backend/internal/study_plan/delivery/http"
)

func registerStudyPlanRoutes(r *echo.Group, h studyplan.Handler) {
	studyPlans := r.Group("/study-plans")

	// collection operations
	http.MapStudyPlanRoutes(studyPlans, h)
}
