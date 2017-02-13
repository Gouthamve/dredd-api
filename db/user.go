package db

import (
	"time"

	"github.com/juju/errors"
)

// User is the user model
type User struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password"`

	Submissions []Submission `json:"submissions"`
	Ticket      Ticket       `json:"ticket"`
}

// BeforeSave is the pre-callback
func (u *User) BeforeSave() error {
	return validateUser(*u)
}

func validateUser(u User) error {
	if u.Name == "" {
		return errors.New("user: name cannot be empty")
	}
	if u.Email == "" {
		return errors.New("user: email cannot be empty")
	}
	return nil
}
