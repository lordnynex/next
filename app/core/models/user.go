package models

import (
	"github.com/sknv/upsale/app/lib/mongo/document"
)

type User struct {
	document.Timestamper `bson:",inline"`

	Email string `json:"email"`
	Name  string `json:"name"`
}
