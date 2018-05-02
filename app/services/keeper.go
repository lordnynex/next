package services

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

type (
	Keeper struct {
		JWTAuth *jwtauth.JWTAuth
	}

	LoginRequest struct {
		Login string
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)

func NewKeeper() *Keeper {
	return &Keeper{JWTAuth: initializers.GetJWTAuth()}
}

func (k *Keeper) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	r.Login = strings.TrimSpace(r.Login)
	if r.Login != "login" {
		return nil, errors.New("user is not authenticated")
	}

	userID := "qwe123"
	_, tokenString, err := k.JWTAuth.Encode(
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
