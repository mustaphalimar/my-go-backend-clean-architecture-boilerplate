package handlers

import (
	exam "github.com/mustaphalimar/prepilotapp-backend/internal/exam"
	examHttp "github.com/mustaphalimar/prepilotapp-backend/internal/exam/delivery/http"
	"github.com/mustaphalimar/prepilotapp-backend/internal/health"
	healthHttp "github.com/mustaphalimar/prepilotapp-backend/internal/health/delivery/http"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
	studyPlanHttp "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan/delivery/http"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/usecases"
)

type Handlers struct {
	Health    health.Handler
	StudyPlan studyplan.Handler
	Exam      exam.Handler
}

func NewHandlers(srv *server.Server, usecases *usecases.Usecases) *Handlers {
	return &Handlers{
		Health:    healthHttp.NewHealthHandler(srv),
		StudyPlan: studyPlanHttp.NewStudyPlanHandler(usecases.StudyPlanUC),
		Exam:      examHttp.NewExamHandler(usecases.ExamUC),
	}
}
