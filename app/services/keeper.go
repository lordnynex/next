package services

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/sknv/upsale/app/lib/net/rpc/proto"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	"github.com/sknv/upsale/app/core/models"
)

const (
	exp = 90 * 24 * time.Hour // Expires in 90 days.
)

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
)

type (
	Keeper struct {
		JWTAuth *jwtauth.JWTAuth
	}

	CreateAuthSessionRequest struct {
		Email string
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

func (k *Keeper) CreateAuthSession(_ context.Context, r *CreateAuthSessionRequest,
) (*proto.Empty, error) {
	email := strings.ToLower(strings.TrimSpace(r.Email))

	// TODO: find user by email, create authsession and send email.
	if email != "user@example.com" {
		return nil, ErrUserDoesNotExist
	}

	authSession := &models.AuthSession{
		ID:        "abc123",
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	log.Print("AuthSession: ", authSession)
	return &proto.Empty{}, nil
}

func (k *Keeper) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	login := strings.TrimSpace(r.Login)
	if login != "login" {
		return nil, errors.New("user is not authenticated")
	}

	userID := "abc123"
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
