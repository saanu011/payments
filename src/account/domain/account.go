package domain

import "time"

type Account struct {
	ID        string    `json:"id" db:"id"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
