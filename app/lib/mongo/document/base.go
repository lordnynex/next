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

func (d *Base) GetID() bson.ObjectId {
	return d.ID
}

func (d *Base) BeforeInsert() {
	d.initID()
}

func (d *Base) initID() {
	if d.ID == "" {
		d.ID = bson.NewObjectId()
	}
}
