package models

import (
	"errors"
	"time"

	"github.com/sknv/next/app/lib/mongo/document"
	"github.com/sknv/upsale/app/lib/utils"
)

const (
	codeExpirationPeriod = 10 * time.Minute
	defaultUserRole      = "client"
)

type User struct {
	document.Timestamper `bson:",inline"`

	Email string `json:"email"`
	Role  string `json:"role"`

	// Fields for passwordless authentication.
	Code          string    `json:"-"`
	CodeCreatedAt time.Time `bson:"code_created_at" json:"-"`
}

func (u *User) BeforeInsert() {
	u.Timestamper.BeforeInsert() // Call super method.

	if u.Role == "" {
		u.Role = defaultUserRole // Set default role if one does not exist yet.
	}
}

func (u *User) Authenticate(code string) error {
	if u.Code == "" {
		return errors.New("code is already used")
	}

	expirationTime := u.CodeCreatedAt.Add(codeExpirationPeriod)
	if time.Now().After(expirationTime) {
		return errors.New("code is expired")
	}

	if u.Code != code {
		return errors.New("code does not match")
	}
	return nil
}

func (u *User) GenerateCode() {
	u.Code = utils.RandomNumbers(6)
	u.CodeCreatedAt = time.Now()
}

func (u *User) LogIn() {
	u.Code = ""
	u.CodeCreatedAt = time.Time{}
}
