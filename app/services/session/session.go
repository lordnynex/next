package session

import (
	"context"

	"github.com/sknv/upsale/app/core/models"

	"github.com/sknv/upsale/app/lib/net/rpc"
)

type (
	Session interface {
		Login(context.Context, *LoginRequest) (*LoginResponse, error)
		IDFromContext(context.Context, *rpc.Empty) (*IDFromContextResponse, error)
		FindOneByID(context.Context, *FindOneByIDRequest) (*FindOneByIDResponse, error)
	}

	LoginRequest struct {
		Login string
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	IDFromContextResponse struct {
		SessionID string
	}

	FindOneByIDRequest struct {
		ID string
	}

	FindOneByIDResponse struct {
		Session *models.Session
	}
)
