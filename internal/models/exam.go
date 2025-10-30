package models

import (
	"time"

	"github.com/google/uuid"
)

type ExamStatus string

const (
	EStatusUpcoming  ExamStatus = "upcoming"
	EStatusPassed    ExamStatus = "passed"
	EStatusCancelled ExamStatus = "cancelled"
)

type Exam struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	Title       string     `gorm:"type:varchar(255);not null" json:"title" validate:"required,min=3,max=255"`
	Description *string    `gorm:"type:text" json:"description,omitempty"`
	ExamDate    time.Time  `gorm:"type:timestamp;not null;column:exam_date" json:"examDate" validate:"required"`
	Status      ExamStatus `gorm:"type:varchar(50);not null;default:'upcoming';column:status" json:"status"`
	PassedAt    *time.Time `gorm:"type:timestamp;column:passed_at" json:"passedAt"`
	Metadata    *Metadata  `gorm:"type:jsonb;serializer:json;column:metadata" json:"metadata"`

	StudyPlanID uuid.UUID `gorm:"type:uuid;not null;unique;column:study_plan_id" json:"studyPlaId" validate:"required"`

	// Relationships
	StudyPlan StudyPlan `gorm:"foreignKey:StudyPlanID" json:"studyPlan,omitempty"`
}
