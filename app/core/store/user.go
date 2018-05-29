package store

import (
	"github.com/globalsign/mgo"

	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/lib/mongo"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (*User) FindOneByID(_ *mgo.Session, id string) (*models.User, error) {
	if id != "abc123" {
		return nil, mgo.ErrNotFound
	}
	return &models.User{ID: "abc123", Email: "user@example.com"}, nil
}

func (*User) FindOneByEmail(_ *mgo.Session, email string) (*models.User, error) {
	if email != "user@example.com" {
		return nil, mgo.ErrNotFound
	}
	return &models.User{ID: "abc123", Email: "user@example.com"}, nil
}

func (u *User) Insert(_ *mgo.Session, user *models.User) error {
	user.ID = "xyz456"
	return nil
}

func (u *User) FindOneOrInsertByEmail(ses *mgo.Session, email string,
) (*models.User, error) {
	user, err := u.FindOneByEmail(ses, email)
	if err == nil {
		return user, nil // Return a user if one exists.
	} else if !mongo.IsErrNotFound(err) {
		return nil, err // Return in case of unknown error.
	}

	// Insert a user if one does not exist yet.
	user = &models.User{Email: email}
	err = u.Insert(ses, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
