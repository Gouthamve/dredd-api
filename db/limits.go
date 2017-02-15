package db

import (
	"time"

	"github.com/gouthamve/dredd"
)

// Limits is the limits model
type Limits struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	dredd.Limits
	ChallengeID uint // Relation
}
