package repositories

import (
	exam "github.com/mustaphalimar/prepilotapp-backend/internal/exam"
	examRepo "github.com/mustaphalimar/prepilotapp-backend/internal/exam/repository"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
	spRepo "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan/repository"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
)

type Repositories struct {
	StudyPlanRepo studyplan.Repository
	ExamRepo      exam.Repository
}

func NewRepositories(server *server.Server) *Repositories {
	studyPlanRepo := spRepo.NewStudyPlanRepository(server)
	examRepo := examRepo.NewExamRepository(server)

	return &Repositories{
		StudyPlanRepo: studyPlanRepo,
		ExamRepo:      examRepo,
	}
}
