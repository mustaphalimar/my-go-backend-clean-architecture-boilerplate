package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	Name          *string    `gorm:"type:varchar(255)" json:"name"`
	Email         string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	ClerkID       string     `gorm:"type:varchar(255);uniqueIndex;not null;column:clerk_id" json:"clerk_id"`
	FirstName     *string    `gorm:"type:varchar(255);column:first_name" json:"first_name"`
	LastName      *string    `gorm:"type:varchar(255);column:last_name" json:"last_name"`
	ImageUrl      *string    `gorm:"type:text;column:image_url" json:"image_url"`
	EmailVerified *bool      `gorm:"type:boolean;default:false;column:email_verified" json:"email_verified"`
	LastSignInAt  *time.Time `gorm:"type:timestamp;column:last_sign_in_at" json:"last_sign_in_at"`
	Banned        *bool      `gorm:"type:boolean;default:false" json:"banned"`
	DeletedAt     *time.Time `gorm:"type:timestamp;column:deleted_at" json:"deleted_at"`

	StudyPlans   []StudyPlan   `gorm:"foreignKey:UserID" json:"study_plans,omitempty"`
	SubjectUsers []SubjectUser `gorm:"foreignKey:UserID" json:"subject_users,omitempty"`
}

type UserProfile struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	UserID uuid.UUID `json:"userId" gorm:"type:uuid;not null"`

	Bio      string     `json:"bio" gorm:"type:text"`
	Avatar   string     `json:"avatar"`
	Timezone string     `json:"timezone" gorm:"default:'UTC'"`
	Language string     `json:"language" gorm:"default:'en'"`
	Birthday *time.Time `json:"birthday"`

	User User `json:"-" gorm:"foreignKey:UserID"`
}

// BeforeCreate hook to ensure UUID is set
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// BeforeCreate hook for UserProfile
func (up *UserProfile) BeforeCreate(tx *gorm.DB) error {
	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	return nil
}

// TableName returns the table name for User model
func (User) TableName() string {
	return "users"
}

// TableName returns the table name for UserProfile model
func (UserProfile) TableName() string {
	return "user_profiles"
}

// FullName returns the user's full name
func (u *User) FullName() string {
	var firstName, lastName string
	if u.FirstName != nil {
		firstName = *u.FirstName
	}
	if u.LastName != nil {
		lastName = *u.LastName
	}
	return firstName + " " + lastName
}

// IsDeleted checks if the user is soft deleted
func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}
