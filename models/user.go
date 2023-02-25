package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
}
