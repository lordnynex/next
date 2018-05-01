package session

import (
	"context"
	"errors"
)

type SessionClient struct{}

func NewSessionClient() Session {
	return &SessionClient{}
}

func (*SessionClient) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	if r.Login != "login" {
		return nil, errors.New("user is not authenticated")
	}
	return &LoginResponse{Token: "123qwe"}, nil
}
