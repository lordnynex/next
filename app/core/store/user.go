package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/lib/mongo"
)

type User struct {
	*Base
}

func NewUser() *User {
	return &User{NewBase("users")}
}

func (u *User) FindOneByID(session *mgo.Session, id string) (*models.User, error) {
	result := &models.User{}
	err := u.Base.FindOneById(session, id, result)
	return result, err
}

func (u *User) FindOneByEmail(session *mgo.Session, email string) (*models.User, error) {
	result := &models.User{}
	err := u.FindOne(session, bson.M{"email": email}, result)
	return result, err
}

func (u *User) FindOneOrInsertByEmail(session *mgo.Session, email string,
) (*models.User, error) {
	user, err := u.FindOneByEmail(session, email)
	if err == nil {
		return user, nil // Return a user if one exists.
	} else if !mongo.IsErrNotFound(err) {
		return nil, err // Return in case of unknown error.
	}

	// Insert a user if one does not exist yet.
	user = &models.User{Email: email}
	err = u.Insert(session, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
