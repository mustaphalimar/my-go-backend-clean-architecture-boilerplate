package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mustaphalimar/prepilotapp-backend/internal/exam"
)

type ExamHandler struct {
	examUC exam.Usecase
}

func NewExamHandler(uc exam.Usecase) exam.Handler {
	return &ExamHandler{
		examUC: uc,
	}
}

func (h *ExamHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *ExamHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
