package services

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/records"
	"github.com/sknv/upsale/app/lib/net/rpc/proto"
)

const (
	exp = 90 * 24 * time.Hour // Expires in 90 days.
)

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
)

type (
	Keeper struct {
		AuthSessions *records.AuthSession
		JWTAuth      *jwtauth.JWTAuth
	}

	CreateAuthSessionRequest struct {
		Email string
	}

	LoginRequest struct {
		AuthSessionID string `json:"auth_session_id"`
	}

	LoginResponse struct {
		AuthToken string `json:"auth_token"`
	}
)

func NewKeeper() *Keeper {
	return &Keeper{
		AuthSessions: records.NewAuthSession(),
		JWTAuth:      initializers.GetJWTAuth(),
	}
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
		UserID:    "abc123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	log.Print("Mail AuthSession: ", authSession)
	return &proto.Empty{}, nil
}

func (k *Keeper) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	authSession, err := k.AuthSessions.FindOneByID(nil, r.AuthSessionID)
	if err != nil {
		return nil, err
	}

	_, tokenString, err := k.JWTAuth.Encode(
		jwtauth.Claims{
			"sub": authSession.UserID,
			"exp": time.Now().Add(exp).Unix(),
		},
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{AuthToken: tokenString}, nil
}
