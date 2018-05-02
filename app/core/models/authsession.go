package models

import "time"

type AuthSession struct {
	ID        string    `json:"id"`
	Email     string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
