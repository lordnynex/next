package models

import (
	"errors"
	"time"

	"github.com/sknv/next/app/lib/mongo/document"
)

const (
	authSessionExpirationPeriod = 15 * time.Minute
)

type AuthSession struct {
	document.Timestamper `bson:",inline"`

	UserID     string    `bson:"user_id" json:"user_id"`
	LoggedInAt time.Time `bson:"logged_in_at" json:"logged_in_at"`
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
