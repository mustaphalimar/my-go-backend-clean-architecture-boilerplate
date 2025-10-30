package models

import (
	"time"

	"github.com/google/uuid"
)

type Analysis struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	AttachmentID uuid.UUID `gorm:"type:uuid;not null;unique;column:attachment_id" json:"attachmentId" validate:"required"`
	Summary      string    `gorm:"type:text;not null;column:summary" json:"summary"`

	KeyTopics []KeyTopics `gorm:"type:jsonb;column:key_topics" json:"keyTopics"`

	Difficulty string `gorm:"type:varchar(20);column:difficulty" json:"difficulty,omitempty"`
	// beginner, intermediate, advanced

	Status string `gorm:"type:varchar(50);not null;default:'pending';column:status" json:"status"`
	// pending, processing, completed, failed

	ErrorMessage *string `gorm:"type:text;column:error_message" json:"errorMessage,omitempty"`

	StudyTime  int     `gorm:"type:integer;column:study_time" json:"studyTime,omitempty"`
	Complexity float64 `gorm:"type:decimal(3,2);column:complexity" json:"complexity,omitempty"`
	// 0.0 to 1.0

	Readability float64 `gorm:"type:decimal(3,2);column:readability" json:"readability,omitempty"`
	// 0.0 to 1.0

	// Relationships
	Insight *Insight `gorm:"foreignKey:AnalysisID" json:"insight,omitempty"`
}

type KeyTopicImportance string

const (
	KTImportanceLow    KeyTopicImportance = "low"
	KTImportanceMedium KeyTopicImportance = "medium"
	KTImportanceHigh   KeyTopicImportance = "high"
)

type KeyTopics struct {
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	Importance      KeyTopicImportance `json:"importance"`
	RelatedChapters int                `json:"relatedChapters"`
}
