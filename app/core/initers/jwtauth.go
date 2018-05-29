package initers

import (
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/cfg"
)

const (
	alg = "HS256"
)

var (
	jwtAuth *jwtauth.JWTAuth
)

func init() {
	jwtAuth = jwtauth.New(alg, []byte(cfg.GetSecretKey()), nil)
}

func GetJWTAuth() *jwtauth.JWTAuth {
	return jwtAuth
}
