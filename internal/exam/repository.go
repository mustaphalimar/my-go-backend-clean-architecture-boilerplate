package exam

import (
	"github.com/google/uuid"
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
)

type Repository interface {
	FindAll() ([]*models.Exam, error)
	FindByID(id uuid.UUID) (*models.Exam, error)
}
