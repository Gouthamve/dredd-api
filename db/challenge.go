package db

import (
	"time"

	"github.com/gouthamve/dredd"
)

// Challenge is the challenge model
type Challenge struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Question  string           `json:"question"`
	Testcases []dredd.Testcase `json:"testcases"`
	Limits    dredd.Limits     `json:"limits"`
}
