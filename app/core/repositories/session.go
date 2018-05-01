package repositories

import (
	"errors"

	"github.com/sknv/upsale/app/core/models"
)

type Session struct{}

func (s *Session) FindOneByID(id string) (*models.Session, error) {
	if id != "123qwe" {
		return nil, errors.New("session is not found")
	}
	return &models.Session{ID: "123qwe", UserID: "456asd"}, nil
}
