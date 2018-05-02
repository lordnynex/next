package records

import (
	"errors"

	"github.com/globalsign/mgo"

	"github.com/sknv/upsale/app/core/models"
)

type AuthSession struct{}

func NewAuthSession() *AuthSession {
	return &AuthSession{}
}

func (*AuthSession) FindOneByID(_ *mgo.Session, id string) (*models.AuthSession, error) {
	if id != "abc123" {
		return nil, errors.New("auth session does not exist")
	}
	return &models.AuthSession{ID: "abc123", UserID: "abc123"}, nil
}
