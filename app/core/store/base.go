package store

import (
	"github.com/sknv/upsale/app/lib/mongo/store"
)

const maxFetchLimit = 25

type Base struct {
	*store.Finder
}

func NewBase(collectionName string) *Base {
	return &Base{
		Finder: &store.Finder{
			Base: &store.Base{
				CollectionName: collectionName,
				MaxFetchLimit:  maxFetchLimit,
			},
		},
	}
}
