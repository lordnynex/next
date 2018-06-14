package initers

import (
	"github.com/globalsign/mgo"

	"github.com/sknv/next/app/core/cfg"
	"github.com/sknv/next/app/lib/mongo"
)

var (
	mongoSession *mgo.Session
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    cfg.GetMongoAddrs(),
		Source:   cfg.GetMongoSource(),
		Database: cfg.GetMongoDatabase(),
		Username: cfg.GetMongoUsername(),
		Password: cfg.GetMongoPassword(),
		Timeout:  cfg.GetMongoTimeout(),
	}
	mongoSession = mongo.MustDial(dialInfo)
}

// GetMongoSession returns a global mgo.Session.
func GetMongoSession() *mgo.Session {
	return mongoSession
}
