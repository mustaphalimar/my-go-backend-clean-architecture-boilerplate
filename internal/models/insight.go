package models

import (
	"time"

	"github.com/google/uuid"
)

type Insight struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`

	AnalysisID uuid.UUID `gorm:"type:uuid;not null;unique;column:analysis_id" json:"analysisId" validate:"required"`
	Overview   string    `gorm:"type:text;not null;column:overview" json:"overview"`

	Chapters []InsightChapter `gorm:"type:jsonb;not null;column:chapters" json:"chapters"`

	KeyConcepts []InsightKeyConcept `gorm:"type:jsonb;column:key_concepts" json:"keyConcepts,omitempty"`

	DependenciesGraph []InsightDependency `gorm:"type:jsonb;column:dependencies_graph" json:"dependenciesGraph"`

	Recommendations []InsightRecommendation `gorm:"type:jsonb;column:recommendations" json:"recommendations"`

	EstimatedStudyTime *int `gorm:"type:integer;column:estimated_study_time" json:"estimatedStudyTime,omitempty"` // minutes

	DifficultyLevel *string `gorm:"type:varchar(20);column:difficulty_level" json:"difficultyLevel,omitempty"` // beginner, intermediate, advanced
}

type InsightChapter struct {
	ChapterNumber      int    `json:"chapterNumber"`
	Title              string `json:"title"`
	Summary            string `json:"summary"`
	PageRange          string `json:"pageRange"`
	EstimatedStudyTime int    `json:"estimatedStudyTime"` // in minutes
}

type InsightKeyConcept struct {
	Concept    string `json:"concept"`
	Definition string `json:"definition"`
	ChapterRef int    `json:"chapterRef,omitempty"`
}
type InsightDependency struct {
	Concept   string   `json:"concept"`
	DependsOn []string `json:"dependsOn"`
}
type InsightRecommendation struct {
	Type    string `json:"type"`    // focus, practice, review, etc.
	Message string `json:"message"` // user-facing tip
}
