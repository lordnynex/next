package session

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	"github.com/sknv/upsale/app/core/repositories"
	"github.com/sknv/upsale/app/lib/net/rpc"
)

const (
	claimSessionID = "sub"
	exp            = 180 * 24 * time.Hour // Expires in 180 days.
)

type SessionClient struct {
	JWTAuth     *jwtauth.JWTAuth
	SessionRepo *repositories.Session
}

func NewSessionClient() Session {
	return &SessionClient{
		JWTAuth:     initializers.NewJWTAuth(),
		SessionRepo: &repositories.Session{},
	}
}

func (c *SessionClient) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	r.Login = strings.TrimSpace(r.Login)
	if r.Login != "login" {
		return nil, errors.New("user is not authenticated")
	}

	sessionID := "123qwe"
	_, tokenString, err := c.JWTAuth.Encode(jwtauth.Claims{claimSessionID: sessionID})
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Token: tokenString}, nil
}

func (c *SessionClient) IDFromContext(ctx context.Context, _ *rpc.Empty,
) (*IDFromContextResponse, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	sessionID, ok := claims[claimSessionID].(string)
	if !ok {
		return nil, errors.New("session id does not exist in jwt claims")
	}
	return &IDFromContextResponse{SessionID: sessionID}, nil
}

func (c *SessionClient) FindOneByID(_ context.Context, r *FindOneByIDRequest,
) (*FindOneByIDResponse, error) {
	session, err := c.SessionRepo.FindOneByID(r.ID)
	if err != nil {
		return nil, err
	}
	return &FindOneByIDResponse{Session: session}, nil
}
