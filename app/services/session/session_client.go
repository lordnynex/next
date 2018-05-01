package session

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
)

const (
	exp = 90 * 24 * time.Hour // Expires in 90 days.
)

type SessionClient struct {
	JWTAuth *jwtauth.JWTAuth
}

func NewSessionClient() Session {
	return &SessionClient{JWTAuth: initializers.NewJWTAuth()}
}

func (c *SessionClient) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	r.Login = strings.TrimSpace(r.Login)
	if r.Login != "login" {
		return nil, errors.New("user is not authenticated")
	}

	userID := "qwe123"
	_, tokenString, err := c.JWTAuth.Encode(
		jwtauth.Claims{
			"sub": userID,
			"exp": time.Now().Add(exp).Unix(),
		},
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Token: tokenString}, nil
}

func (c *SessionClient) GetUserID(_ context.Context, r *GetUserIDRequest,
) (*GetUserIDResponse, error) {
	_, claims, _ := jwtauth.FromContext(r.Context)
	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("session id does not exist in jwt claims")
	}
	return &GetUserIDResponse{UserID: userID}, nil
}
