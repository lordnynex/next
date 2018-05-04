package services

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/sknv/upsale/app/core/initializers"
	"github.com/sknv/upsale/app/core/mailers"
	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/records"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
	"github.com/sknv/upsale/app/lib/net/rpc/proto"
)

const (
	exp = 90 * 24 * time.Hour // Expires in 90 days.
)

type (
	AuthKeeper struct {
		JWTAuth      *jwtauth.JWTAuth
		LoginMailer  *mailers.Login
		AuthSessions *records.AuthSession
		Users        *records.User
	}

	CreateAuthSessionRequest struct {
		Email string `json:"email"`
	}

	LoginResponse struct {
		AuthToken string `json:"auth_token"`
	}
)

func NewAuthKeeper() *AuthKeeper {
	return &AuthKeeper{
		JWTAuth:      initializers.GetJWTAuth(),
		LoginMailer:  mailers.NewLogin(),
		AuthSessions: records.NewAuthSession(),
		Users:        records.NewUser(),
	}
}

// CreateAuthSession creates a user account if one does not exist yet and stores an auth session.
func (a *AuthKeeper) CreateAuthSession(_ context.Context, r *CreateAuthSessionRequest,
) (*proto.Empty, error) {
	if err := r.Validate(); err != nil {
		return nil, &xhttp.ErrHttpStatus{Err: err, Status: http.StatusUnprocessableEntity}
	}

	// Create an account if one does not exist yet.
	user, err := a.Users.FindOneOrInsertByEmail(nil, r.Email)
	if err != nil {
		return nil, &xhttp.ErrHttpStatus{Err: err, Status: http.StatusInternalServerError}
	}

	authSession := &models.AuthSession{UserID: user.ID}
	if err := a.AuthSessions.Insert(nil, authSession); err != nil {
		return nil, &xhttp.ErrHttpStatus{Err: err, Status: http.StatusInternalServerError}
	}

	go a.LoginMailer.Deliver(authSession.ID, user.Email) // Deliver later.
	return &proto.Empty{}, nil
}

// Login validates an auth session and in case of a successful validation returns an auth token.
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
		jwtauth.Claims{"sub": authSession.UserID, "exp": time.Now().Add(exp).Unix()},
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{AuthToken: tokenString}, nil
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

func (r *CreateAuthSessionRequest) Validate() error {
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Email, validation.Required, is.Email),
	)
}
