package store

import (
	"github.com/sknv/next/app/lib/mongo/document"
)

func doBeforeInsertIfNeeded(doc interface{}) {
	if beforeInserter, ok := doc.(document.IBeforeInserter); ok {
		beforeInserter.BeforeInsert()
	}
}

func doBeforeRemoveIfNeeded(doc interface{}) {
	if beforeRemover, ok := doc.(document.IBeforeRemover); ok {
		beforeRemover.BeforeRemove()
	}
}

func doBeforeSaveIfNeeded(doc interface{}) {
	if beforeSaver, ok := doc.(document.IBeforeSaver); ok {
		beforeSaver.BeforeSave()
	}
}

func doBeforeUpdateIfNeeded(doc interface{}) {
	if beforeUpdater, ok := doc.(document.IBeforeUpdater); ok {
		beforeUpdater.BeforeUpdate()
	}
}

func doAfterInsertIfNeeded(doc interface{}) {
	if afterInserter, ok := doc.(document.IAfterInserter); ok {
		afterInserter.AfterInsert()
	}
}

func doAfterRemoveIfNeeded(doc interface{}) {
	if afterRemover, ok := doc.(document.IAfterRemover); ok {
		afterRemover.AfterRemove()
	}
}

func doAfterSaveIfNeeded(doc interface{}) {
	if afterSaver, ok := doc.(document.IAfterSaver); ok {
		afterSaver.AfterSave()
	}
}

func doAfterUpdateIfNeeded(doc interface{}) {
	if afterUpdater, ok := doc.(document.IAfterUpdater); ok {
		afterUpdater.AfterUpdate()
	}
}
