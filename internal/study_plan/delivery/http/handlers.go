package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
)

type StudyPlanHandler struct {
	studyPlanUC studyplan.Usecase
}

func NewStudyPlanHandler(uc studyplan.Usecase) studyplan.Handler {
	return &StudyPlanHandler{
		studyPlanUC: uc,
	}
}

func (h *StudyPlanHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "Hello from study plan handler",
		})
	}
}

func (h *StudyPlanHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "Hello from study plan handler",
		})
	}
}
