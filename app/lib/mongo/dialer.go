package mongo

import (
	"github.com/globalsign/mgo"
)

func MustDial(dialInfo *mgo.DialInfo) *mgo.Session {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	return session
}
