package models

import (
	"time"

	"github.com/google/uuid"
)

type Topic struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	Name        string  `gorm:"type:varchar(255);not null;unique" json:"name"` // unique to prevent duplicates
	Description *string `gorm:"type:text" json:"description"`
	Icon        *string `gorm:"type:varchar(100)" json:"icon"`                      // Icon identifier (e.g., "science", "math")
	Color       *string `gorm:"type:varchar(7)" json:"color"`                       // Hex color code for UI theming
	IsActive    bool    `gorm:"type:boolean;not null;default:true" json:"isActive"` // For admin control
	SortOrder   int     `gorm:"type:integer;not null;default:0" json:"sortOrder"`   // For ordering topics

	// Relationships
	Subjects []Subject `gorm:"foreignKey:TopicID" json:"subjects,omitempty"`
}

// TableName returns the table name for Topic model
func (Topic) TableName() string {
	return "topics"
}
