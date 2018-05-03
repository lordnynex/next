package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	"github.com/sknv/upsale/app/core/mailers"
	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/records"
	"github.com/sknv/upsale/app/lib/net/rpc/proto"
)

const (
	exp = 90 * 24 * time.Hour // Expires in 90 days.
)

type (
	AuthKeeper struct {
		AuthSessions *records.AuthSession
		JWTAuth      *jwtauth.JWTAuth
		LoginMailer  *mailers.Login
	}

	CreateAuthSessionRequest struct {
		Email string
	}

	LoginResponse struct {
		AuthToken string `json:"auth_token"`
	}
)

func NewAuthKeeper() *AuthKeeper {
	return &AuthKeeper{
		AuthSessions: records.NewAuthSession(),
		JWTAuth:      initializers.GetJWTAuth(),
		LoginMailer:  mailers.NewLogin(),
	}
}

func (a *AuthKeeper) CreateAuthSession(_ context.Context, r *CreateAuthSessionRequest,
) (*proto.Empty, error) {
	email := strings.ToLower(strings.TrimSpace(r.Email))
	// TODO: find user by email, create authsession and send email.
	if email != "user@example.com" {
		return nil, errors.New("user does not exist")
	}

	authSession := &models.AuthSession{
		ID:        "abc123",
		UserID:    "abc123",
		CreatedAt: time.Now(),
	}
	go a.LoginMailer.Deliver(authSession.ID, email)
	return &proto.Empty{}, nil
}

func (a *AuthKeeper) Login(_ context.Context, authSessionID string) (*LoginResponse, error) {
	authSession, err := a.AuthSessions.FindOneByID(nil, authSessionID)
	if err != nil {
		return nil, err
	}

	if err := authSession.Validate(); err != nil {
		return nil, err
	}

	authSession.LogIn()
	if err := a.AuthSessions.UpdateDoc(nil, authSession); err != nil {
		return nil, err
	}

	_, tokenString, err := a.JWTAuth.Encode(
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
