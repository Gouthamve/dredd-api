package db

import (
	"time"

	"github.com/gouthamve/dredd"
)

// Testcase is the limits model
type Testcase struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	dredd.Testcase
	ChallengeID uint // Relation
}
