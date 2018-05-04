package models

import (
	"errors"
	"time"
)

const (
	authSessionExpirationPeriod = 15 * time.Minute
)

type AuthSession struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	LoggedInAt time.Time `json:"logged_in_at"`
}

func (a *AuthSession) Validate() error {
	if !a.LoggedInAt.IsZero() {
		return errors.New("auth session is already used")
	}

	expirationTime := a.CreatedAt.Add(authSessionExpirationPeriod)
	if time.Now().After(expirationTime) {
		return errors.New("auth session is expired")
	}
	return nil
}

func (a *AuthSession) LogIn() {
	a.LoggedInAt = time.Now()
}
