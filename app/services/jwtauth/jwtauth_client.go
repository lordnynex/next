package jwtauth

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth"
)

const (
	alg = "HS256"
	exp = 180 * 24 * time.Hour // Expires in 180 days.
)

type JWTAuthClient struct {
	JWTAuth *jwtauth.JWTAuth
}

func NewJWTAuthClient(secretKey string) JWTAuth {
	return &JWTAuthClient{JWTAuth: jwtauth.New(alg, []byte(secretKey), nil)}
}

func (j *JWTAuthClient) Encode(_ context.Context, r *EncodeRequest) (*EncodeResponse, error) {
	r.Payload["exp"] = time.Now().Add(exp).Unix()

	_, tokenString, err := j.JWTAuth.Encode(r.Payload)
	if err != nil {
		return nil, err
	}
	return &EncodeResponse{Token: tokenString}, nil
}
