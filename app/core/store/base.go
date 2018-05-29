package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/sknv/upsale/app/lib/mongo/store"
)

const maxLimit = 25

type Base struct {
	*store.Base
}

func NewBase(collectionName string) *Base {
	return &Base{
		&store.Base{
			CollectionName: collectionName,
			MaxLimit:       maxLimit,
		},
	}
}

func (r *Base) FindAll(
	session *mgo.Session, query bson.M, sort []string, result interface{},
) error {
	qry := r.Find(session, query)
	if len(sort) > 0 {
		qry = qry.Sort(sort...)
	}
	return qry.All(result)
}

func (r *Base) FindOne(session *mgo.Session, query bson.M, result interface{}) error {
	qry := r.Find(session, query)
	return qry.One(result)
}

func (r *Base) FindOneById(session *mgo.Session, id string, result interface{}) error {
	qry := r.Find(session, bson.M{"_id": bson.ObjectIdHex(id)})
	return qry.One(result)
}

func (r *Base) FindPage(
	session *mgo.Session, query bson.M, params store.PagingParams, result interface{},
) error {
	qry := r.Base.FindPage(session, query, params)
	return qry.All(result)
}
