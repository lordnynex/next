package store

import (
	"github.com/globalsign/mgo"

	"github.com/sknv/upsale/app/core/models"
)

type AuthSession struct {
	*Base
}

func NewAuthSession() *AuthSession {
	return &AuthSession{NewBase("authsessions")}
}

func (a *AuthSession) FindOneByID(session *mgo.Session, id string,
) (*models.AuthSession, error) {
	result := &models.AuthSession{}
	err := a.Base.FindOneById(session, id, result)
	return result, err
}
