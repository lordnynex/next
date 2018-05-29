package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Finder struct {
	*Base
}

func (r *Finder) FindAll(
	session *mgo.Session, query bson.M, sort []string, result interface{},
) error {
	qry := r.Find(session, query)
	if len(sort) > 0 {
		qry = qry.Sort(sort...)
	}
	return qry.All(result)
}

func (r *Finder) FindOne(session *mgo.Session, query bson.M, result interface{}) error {
	qry := r.Find(session, query)
	return qry.One(result)
}

func (r *Finder) FindOneById(session *mgo.Session, id string, result interface{}) error {
	qry := r.Find(session, bson.M{"_id": bson.ObjectIdHex(id)})
	return qry.One(result)
}

func (r *Finder) FindPage(
	session *mgo.Session, query bson.M, params PagingParams, result interface{},
) error {
	qry := r.Base.FindPage(session, query, params)
	return qry.All(result)
}
