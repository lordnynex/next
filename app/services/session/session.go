package session

import (
	"context"
)

const (
	ClaimSessionID = "sub"
)

type (
	Session interface {
		Login(context.Context, *LoginRequest) (*LoginResponse, error)
	}

	LoginRequest struct {
		Login string
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
