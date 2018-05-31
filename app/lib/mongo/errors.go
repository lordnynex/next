package mongo

import (
	"github.com/globalsign/mgo"
)

func IsErrNotFound(err error) bool {
	return err == mgo.ErrNotFound
}
