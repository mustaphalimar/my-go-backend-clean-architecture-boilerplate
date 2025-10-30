package studyplan

import (
	"github.com/google/uuid"
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
)

type Repository interface {
	FindAll() ([]*models.StudyPlan, error)
	FindByID(id uuid.UUID) (*models.StudyPlan, error)
}
