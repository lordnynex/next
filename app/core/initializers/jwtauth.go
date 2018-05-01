package initializers

import (
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/cfg"
)

const (
	alg = "HS256"
)

func NewJWTAuth() *jwtauth.JWTAuth {
	return jwtauth.New(alg, []byte(cfg.GetSecretKey()), nil)
}
