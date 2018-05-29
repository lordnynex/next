package document

type IBeforeInserter interface {
	BeforeInsert()
}

type IBeforeRemover interface {
	BeforeRemove()
}

type IBeforeSaver interface {
	BeforeSave()
}

type IBeforeUpdater interface {
	BeforeUpdate()
}

type IAfterInserter interface {
	AfterInsert()
}

type IAfterRemover interface {
	AfterRemove()
}

type IAfterSaver interface {
	AfterSave()
}

type IAfterUpdater interface {
	AfterUpdate()
}
