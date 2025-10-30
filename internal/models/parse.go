package models

import (
	"time"

	"github.com/google/uuid"
)

type Parse struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	AttachmentID uuid.UUID `gorm:"type:uuid;not null;unique;column:attachment_id" json:"attachmentId" validate:"required"`

	ContentKey *string `gorm:"type:text;column:content_key" json:"contentKey"`

	Status string `gorm:"type:varchar(50);not null;default:'pending';column:status" json:"status"` // pending, processing, completed, failed

	ErrorMessage *string `gorm:"type:text;column:error_message" json:"errorMessage,omitempty"`

	// Metadata
	PageCount *int `gorm:"type:integer;column:page_count" json:"pageCount,omitempty"`
	WordCount *int `gorm:"type:integer;column:word_count" json:"wordCount,omitempty"`
	CharCount *int `gorm:"type:integer;column:char_count" json:"charCount,omitempty"`
}
