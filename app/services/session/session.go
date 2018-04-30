package session

import (
	"context"
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
