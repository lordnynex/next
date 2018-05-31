package models

import (
	"github.com/sknv/next/app/lib/mongo/document"
)

const (
	defaultUserRole = "client"
)

type User struct {
	document.Timestamper `bson:",inline"`

	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func (u *User) BeforeInsert() {
	if u.Role == "" {
		u.Role = defaultUserRole // Set default role if one does not exist yet.
	}
}
