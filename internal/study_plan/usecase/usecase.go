package usecase

import (
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
)

type StudyPlanUC struct {
	server *server.Server
	repo   studyplan.Repository
}

func NewStudyPlanUC(srv *server.Server, repo studyplan.Repository) studyplan.Usecase {
	return &StudyPlanUC{
		repo: repo,
	}
}

func (uc *StudyPlanUC) GetAll() ([]*models.StudyPlan, error) {
	return nil, nil
}

func (uc *StudyPlanUC) GetByID(id int64) (*models.StudyPlan, error) {
	return nil, nil
}
