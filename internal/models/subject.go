package models

import (
	"time"

	"github.com/google/uuid"
)

type Subject struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	Name            string    `gorm:"type:varchar(255);not null" json:"name"` // removed unique constraint since topics can have same subject names
	Description     *string   `gorm:"type:text" json:"description"`
	StudyPlansCount int       `gorm:"type:integer;not null;default:0;column:study_plans_count" json:"studyPlansCount"`
	TopicID         uuid.UUID `gorm:"type:uuid;not null;column:topic_id" json:"topicId"`
	SortOrder       int       `gorm:"type:integer;not null;default:0" json:"sortOrder"`   // For ordering subjects within topic
	IsActive        bool      `gorm:"type:boolean;not null;default:true" json:"isActive"` // For admin control

	// Relationships
	Topic        Topic         `gorm:"foreignKey:TopicID;references:ID" json:"topic,omitempty"`
	StudyPlans   []StudyPlan   `gorm:"foreignKey:SubjectID" json:"study_plans,omitempty"`
	SubjectUsers []SubjectUser `gorm:"foreignKey:SubjectID" json:"subject_users,omitempty"`
}

type SubjectUser struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_subject_user_unique,unique;column:user_id" json:"userId"`
	SubjectID uuid.UUID `gorm:"type:uuid;not null;index:idx_subject_user_unique,unique;column:subject_id" json:"subjectId"`
	Metadata  *Metadata `gorm:"type:jsonb;serializer:json;column:metadata" json:"metadata"`

	// Relationships
	User    User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Subject Subject `gorm:"foreignKey:SubjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"subject,omitempty"`
}
