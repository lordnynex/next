package jwtauth

import (
	"context"

	"github.com/go-chi/jwtauth"
)

type (
	JWTAuth interface {
		Encode(context.Context, *EncodeRequest) (*EncodeResponse, error)
	}

	EncodeRequest struct {
		Payload jwtauth.Claims
	}

	EncodeResponse struct {
		Token string `json:"token"`
	}
)
