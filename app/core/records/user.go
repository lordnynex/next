package records

import (
	"errors"

	"github.com/globalsign/mgo"

	"github.com/sknv/upsale/app/core/models"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (*User) FindOneByID(_ *mgo.Session, id string) (*models.User, error) {
	if id != "abc123" {
		return nil, errors.New("user does not exist")
	}
	return &models.User{ID: "abc123", Email: "user@example.com"}, nil
}
