package document

import (
	"time"
)

type Timestamper struct {
	Base `bson:",inline"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (t *Timestamper) BeforeInsert() {
	t.Base.BeforeInsert() // Call super method.

	t.initTimestamps()
}

func (t *Timestamper) BeforeSave() {
	t.updateTimestamps()
}

func (t *Timestamper) initTimestamps() {
	t.CreatedAt = time.Now()
}

func (t *Timestamper) updateTimestamps() {
	t.UpdatedAt = time.Now()
}
