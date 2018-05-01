package session

import (
	"context"
)

type (
	Session interface {
		Login(context.Context, *LoginRequest) (*LoginResponse, error)
		GetUserID(context.Context, *GetUserIDRequest) (*GetUserIDResponse, error)
	}

	LoginRequest struct {
		Login string
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	GetUserIDRequest struct {
		Context context.Context
	}

	GetUserIDResponse struct {
		UserID string
	}
)
