package store

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/sknv/upsale/app/lib/mongo/document"
)

type (
	Base struct {
		CollectionName string
		MaxFetchLimit  int
	}

	PagingParams struct {
		Limit int
		Skip  int
		Sort  []string
	}
)

func (b *Base) CollectionForDb(db *mgo.Database) *mgo.Collection {
	return db.C(b.CollectionName)
}

func (b *Base) CollectionForSession(session *mgo.Session) *mgo.Collection {
	db := session.DB("")
	return b.CollectionForDb(db)
}

func (b *Base) Find(session *mgo.Session, query bson.M) *mgo.Query {
	c := b.CollectionForSession(session)
	return c.Find(query)
}

func (b *Base) FindPage(session *mgo.Session, query bson.M, params PagingParams,
) *mgo.Query {
	qry := b.Find(session, query)

	// Set limit and skip params.
	limit := b.MaxFetchLimit
	if params.Limit > 0 && params.Limit < limit {
		limit = params.Limit // Restrict fetching limit.
	}
	qry = qry.Limit(limit)

	if params.Skip > 0 {
		qry = qry.Skip(params.Skip)
	}

	// Sort query.
	if len(params.Sort) > 0 {
		qry = qry.Sort(params.Sort...)
	}
	return qry
}

func (b *Base) Insert(session *mgo.Session, doc interface{}) error {
	// Before callbacks section.
	doBeforeInsertIfNeeded(doc)
	doBeforeSaveIfNeeded(doc)

	col := b.CollectionForSession(session)
	if err := col.Insert(doc); err != nil {
		return err
	}

	// After callbacks section.
	doAfterInsertIfNeeded(doc)
	doAfterSaveIfNeeded(doc)

	return nil
}

func (b *Base) Update(session *mgo.Session, selector interface{}, update interface{},
) error {
	// Before callbacks section.
	doBeforeUpdateIfNeeded(update)
	doBeforeSaveIfNeeded(update)

	col := b.CollectionForSession(session)
	if err := col.Update(selector, update); err != nil {
		return err
	}

	// After callbacks section.
	doAfterUpdateIfNeeded(update)
	doAfterSaveIfNeeded(update)

	return nil
}

func (b *Base) UpdateDoc(session *mgo.Session, doc document.IIdentifier) error {
	return b.Update(session, bson.M{"_id": doc.GetID()}, doc)
}

func (b *Base) Remove(session *mgo.Session, selector interface{}) error {
	col := b.CollectionForSession(session)
	return col.Remove(selector)
}

func (b *Base) RemoveDoc(session *mgo.Session, doc document.IIdentifier) error {
	// Before callbacks section.
	doBeforeRemoveIfNeeded(doc)

	if err := b.Remove(session, bson.M{"_id": doc.GetID()}); err != nil {
		return err
	}

	// After callbacks section.
	doAfterRemoveIfNeeded(doc)

	return nil
}
