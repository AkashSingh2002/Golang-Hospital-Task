package models

import (
	"time"
)

// Patient represents the patient model in the database.
type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Gender    string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}