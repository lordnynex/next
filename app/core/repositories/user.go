package repositories

import (
	"errors"

	"github.com/sknv/upsale/app/core/models"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (*User) FindOneByID(id string) (*models.User, error) {
	if id != "qwe123" {
		return nil, errors.New("user does not exist")
	}
	return &models.User{ID: "qwe123", Username: "login"}, nil
}
