package store

import (
	"time"

	"github.com/globalsign/mgo"

	"github.com/sknv/upsale/app/core/models"
)

type AuthSession struct{}

func NewAuthSession() *AuthSession {
	return &AuthSession{}
}

func (*AuthSession) FindOneByID(_ *mgo.Session, id string) (*models.AuthSession, error) {
	if id != "abc123" {
		return nil, mgo.ErrNotFound
	}
	return &models.AuthSession{
		ID:        "abc123",
		UserID:    "abc123",
		CreatedAt: time.Now().Add(-10 * time.Minute),
	}, nil
}

func (*AuthSession) Insert(_ *mgo.Session, authSession *models.AuthSession) error {
	authSession.ID = "abc123"
	return nil
}

func (*AuthSession) UpdateDoc(_ *mgo.Session, authSession *models.AuthSession) error {
	return nil
}
