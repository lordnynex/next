package initers

import (
	"github.com/globalsign/mgo"

	"github.com/sknv/upsale/app/core/cfg"
	"github.com/sknv/upsale/app/lib/mongo"
)

var (
	mongoSession *mgo.Session
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    cfg.GetMongoAddrs(),
		Database: cfg.GetMongoDatabase(),
		Username: cfg.GetMongoUsername(),
		Password: cfg.GetMongoPassword(),
		Timeout:  cfg.GetMongoTimeout(),
	}
	mongoSession = mongo.MustDial(dialInfo)
}

// GetMongoSession returns a copy of global mgo.Session.
// Remember to call "defer session.Close()" after calling this function.
func GetMongoSession() *mgo.Session {
	return mongoSession.Copy()
}

func CloseMongoSession() {
	mongoSession.Close()
}
