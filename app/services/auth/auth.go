package auth

import (
	"context"

	"github.com/sknv/upsale/app/core/models"
)

type (
	Auth interface {
		Login(context.Context, *LoginRequest) (*LoginResponse, error)
		GetCurrentUser(context.Context, *GetCurrentUserRequest) (*CurrentUserResponse, error)
	}

	LoginRequest struct {
		Login string
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	GetCurrentUserRequest struct {
		Context context.Context
	}

	CurrentUserResponse struct {
		User *models.User
	}
)
