package session

import (
	"context"
	"errors"
	"strings"

	"github.com/go-chi/jwtauth"
	xjwtauth "github.com/sknv/upsale/app/services/jwtauth"
)

type SessionClient struct {
	JWTAuthClient xjwtauth.JWTAuth
}

func NewSessionClient() Session {
	return &SessionClient{JWTAuthClient: xjwtauth.NewJWTAuthClient()}
}

func (s *SessionClient) Login(ctx context.Context, r *LoginRequest) (*LoginResponse, error) {
	r.Login = strings.TrimSpace(r.Login)
	if r.Login != "login" {
		return nil, errors.New("user is not authenticated")
	}

	sessionID := "123qwe"
	encodeResponse, err := s.JWTAuthClient.Encode(
		ctx, &xjwtauth.EncodeRequest{
			Payload: jwtauth.Claims{"sub": sessionID},
		},
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Token: encodeResponse.Token}, nil
}
