package document

import (
	"github.com/globalsign/mgo/bson"
)

type (
	IIdentifier interface {
		GetID() bson.ObjectId
	}

	Base struct {
		ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	}
)

func (b *Base) GetID() bson.ObjectId {
	return b.ID
}

func (b *Base) BeforeInsert() {
	b.initID()
}

func (b *Base) initID() {
	if b.ID == "" {
		b.ID = bson.NewObjectId() // Generate ID if one does not exist yet.
	}
}
