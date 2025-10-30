package studyplan

import "github.com/mustaphalimar/prepilotapp-backend/internal/models"

type Usecase interface {
	GetAll() ([]*models.StudyPlan, error)
	GetByID(id int64) (*models.StudyPlan, error)
}
