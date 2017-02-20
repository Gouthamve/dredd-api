package db

import "time"

// Submission is the model for the file submission
type Submission struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	FileURL string `json:"file_url"`
	Status  string `json:"status"`

	UserID uint // Relation

	Challenge   Challenge `json:"challenge"`
	ChallengeID uint      // Relation
}
