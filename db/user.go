package db

import (
	"time"

	"golang.org/x/crypto/bcrypt"

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
	Ticket      Ticket       `json:"ticket,omitempty"`

	IsAdmin bool `json:"-"`
}

// BeforeSave is the pre-callback
func (u *User) BeforeSave() error {
	return validateUser(*u)
}

// BeforeCreate is the before create callback
func (u *User) BeforeCreate() error {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	u.Password = string(pass)
	return nil
}

func validateUser(u User) error {
	if u.Name == "" {
		return errors.New("user: name cannot be empty")
	}
	if u.Email == "" {
		return errors.New("user: email cannot be empty")
	}

	if u.Password == "" {
		return errors.New("user: password cannot be empty")
	}

	return nil
}

// AfterFind is the after find callback
func (u *User) AfterFind() error {
	u.Password = ""

	return nil
}
