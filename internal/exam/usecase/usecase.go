package usecase

import (
	"github.com/mustaphalimar/prepilotapp-backend/internal/exam"
	"github.com/mustaphalimar/prepilotapp-backend/internal/models"
)

type ExamUC struct {
	repo exam.Repository
}

func NewExamUC(repo exam.Repository) exam.Usecase {
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
