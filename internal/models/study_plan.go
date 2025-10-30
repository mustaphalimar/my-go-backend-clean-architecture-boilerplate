package models

import (
	"time"

	"github.com/google/uuid"
)

type StudyPlanStatus string

const (
	SPStatusActive    StudyPlanStatus = "active"
	SPStatusCompleted StudyPlanStatus = "completed"
	SPStatusArchived  StudyPlanStatus = "archived"
)

type StudyPlan struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	Title       string          `gorm:"type:varchar(255);not null" json:"title" validate:"required,min=3,max=255"`
	Description *string         `gorm:"type:text" json:"description,omitempty"`
	Status      StudyPlanStatus `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
	SubjectID   uuid.UUID       `gorm:"type:uuid;column:subject_id" json:"subject_id"`
	UserID      uuid.UUID       `gorm:"type:uuid;not null;column:user_id" json:"user_id"`

	// Relationships
	User    User    `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Subject Subject `gorm:"foreignKey:SubjectID;references:ID" json:"subject,omitempty"`
	Exam    *Exam   `gorm:"foreignKey:StudyPlanID;references:ID" json:"exam,omitempty"`

	Metadata *Metadata `gorm:"type:jsonb;serializer:json;column:metadata" json:"metadata"`

	Attachments *[]StudyPlanAttachment `gorm:"foreignKey:StudyPlanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"attachments,omitempty"`
}

type StudyPlanAttachment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	Name        string  `gorm:"type:text;not null;column:name" json:"name"`
	UploadedBy  string  `gorm:"type:text;not null;column:uploaded_by" json:"uploadedBy"`
	DownloadKey string  `gorm:"type:text;not null;column:download_key" json:"downloadKey"`
	FileSize    *int64  `gorm:"type:bigint;column:file_size" json:"fileSize"`
	MimeType    *string `gorm:"type:text;column:mime_type" json:"mimeType"`
	Parsed      bool    `gorm:"type:boolean;not null;default:false;column:parsed" json:"parsed"`
	Analyzed    bool    `gorm:"type:boolean;not null;default:false;column:analyzed" json:"analyzed"`

	StudyPlanID uuid.UUID `gorm:"type:uuid;not null;unique;column:study_plan_id" json:"studyPlanId" validate:"required"`

	// Relationships
	StudyPlan StudyPlan `gorm:"foreignKey:StudyPlanID" json:"studyPlan,omitempty"`
	Parse     *Parse    `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE" json:"parse,omitempty"`
	Analysis  *Analysis `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE" json:"analysis,omitempty"`
}
