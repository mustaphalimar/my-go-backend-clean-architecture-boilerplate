package exam

import "github.com/mustaphalimar/prepilotapp-backend/internal/models"

type Usecase interface {
	GetAll() ([]*models.Exam, error)
	GetByID(id int64) (*models.Exam, error)
}
