package jwtauth

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
)

const (
	exp = 180 * 24 * time.Hour // Expires in 180 days.
)

type JWTAuthClient struct {
	JWTAuth *jwtauth.JWTAuth
}

func NewJWTAuthClient() JWTAuth {
	return &JWTAuthClient{JWTAuth: initializers.NewJWTAuth()}
}

func (c *JWTAuthClient) Encode(_ context.Context, r *EncodeRequest) (*EncodeResponse, error) {
	r.Payload["exp"] = time.Now().Add(exp).Unix()

	_, tokenString, err := c.JWTAuth.Encode(r.Payload)
	if err != nil {
		return nil, err
	}
	return &EncodeResponse{Token: tokenString}, nil
}
