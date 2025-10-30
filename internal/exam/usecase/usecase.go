package usecase

import (
	"github.com/mustaphalimar/prepilotapp-backend/internal/exam"
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/server"
)

type ExamUC struct {
	server *server.Server
	repo   exam.Repository
}

func NewExamUC(srv *server.Server, repo exam.Repository) exam.Usecase {
	return &ExamUC{
		repo: repo,
	}
}

func (uc *ExamUC) GetAll() ([]*models.Exam, error) {
	return nil, nil
}
func (uc *ExamUC) GetByID(id int64) (*models.Exam, error) {
	return nil, nil
}
