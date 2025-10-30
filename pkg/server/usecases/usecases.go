package usecases

import (
	exam "github.com/mustaphalimar/prepilotapp-backend/internal/exam"
	examUC "github.com/mustaphalimar/prepilotapp-backend/internal/exam/usecase"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
	studyPlanUC "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan/usecase"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server/repositories"
)

type Usecases struct {
	StudyPlanUC studyplan.Usecase
	ExamUC      exam.Usecase
}

func NewUsecases(repos *repositories.Repositories) *Usecases {

	return &Usecases{
		StudyPlanUC: studyPlanUC.NewStudyPlanUC(repos.StudyPlanRepo),
		ExamUC:      examUC.NewExamUC(repos.ExamRepo),
	}
}
