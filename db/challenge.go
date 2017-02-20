package db

import (
	"errors"
	"time"
)

// Challenge is the challenge model
type Challenge struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Question  string     `json:"question"`
	Testcases []Testcase `json:"testcases"`
	Limits    Limits     `json:"limits"`

	Submissions []Submission `json:"submissions,omitempty"`
}

// BeforeSave is the pre-callback
func (ch *Challenge) BeforeSave() error {
	return validateChallenge(*ch)
}

func validateChallenge(ch Challenge) error {
	if ch.Question == "" {
		return errors.New("challenge: question cannot be empty")
	}
	if ch.Testcases == nil || len(ch.Testcases) == 0 {
		return errors.New("challenge: testcases cannot be empty")
	}

	if ch.Limits == (Limits{}) {
		return errors.New("challenge: limits cannot be empty")
	}

	return nil
}
