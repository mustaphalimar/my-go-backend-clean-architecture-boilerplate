package repository

import (
	"github.com/google/uuid"
	"github.com/mustaphalimar/prepilotapp-backend/internal/exam"
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
)

type ExamRepository struct {
	server *server.Server
}

func NewExamRepository(server *server.Server) exam.Repository {
	return &ExamRepository{
		server: server,
	}
}

func (r *ExamRepository) FindAll() ([]*models.Exam, error) {
	return nil, nil
}

func (r *ExamRepository) FindByID(id uuid.UUID) (*models.Exam, error) {
	return nil, nil
}
