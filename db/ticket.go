package db

import "time"

// Ticket is the ticket model
type Ticket struct {
	// The compulsory fields
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Status string `json:"status"`
	URL    string `json:"url"`
}
