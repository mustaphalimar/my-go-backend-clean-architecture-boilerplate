package repository

import (
	"github.com/google/uuid"
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
	studyplan "github.com/mustaphalimar/prepilotapp-backend/internal/study_plan"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
)

type StudyPlanRepository struct {
	server *server.Server
}

func NewStudyPlanRepository(server *server.Server) studyplan.Repository {
	return &StudyPlanRepository{
		server: server,
	}
}

func (r *StudyPlanRepository) FindAll() ([]*models.StudyPlan, error) {

	return nil, nil
}

func (r *StudyPlanRepository) FindByID(id uuid.UUID) (*models.StudyPlan, error) {
	return nil, nil
}
