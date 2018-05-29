package services

import (
	"net/http"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-chi/jwtauth"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/sknv/upsale/app/core/initers"
	"github.com/sknv/upsale/app/core/mailers"
	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/store"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
)

const (
	exp = 90 * 24 * time.Hour // Expires in 90 days.
)

type (
	AuthKeeper struct {
		JWTAuth      *jwtauth.JWTAuth
		LoginMailer  *mailers.Login
		AuthSessions *store.AuthSession
		Users        *store.User
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
		JWTAuth:      initers.GetJWTAuth(),
		LoginMailer:  mailers.NewLogin(),
		AuthSessions: store.NewAuthSession(),
		Users:        store.NewUser(),
	}
}

// CreateAuthSession creates a user account if one does not exist yet and stores an auth session.
func (a *AuthKeeper) CreateAuthSession(
	mongoSession *mgo.Session, r *CreateAuthSessionRequest,
) error {
	if err := r.Validate(); err != nil {
		return &xhttp.ErrHttpStatus{Err: err, Status: http.StatusUnprocessableEntity}
	}

	// Create an account if one does not exist yet.
	user, err := a.Users.FindOneOrInsertByEmail(mongoSession, r.Email)
	if err != nil {
		return &xhttp.ErrHttpStatus{Err: err, Status: http.StatusInternalServerError}
	}

	authSession := &models.AuthSession{UserID: user.ID.Hex()}
	if err := a.AuthSessions.Insert(mongoSession, authSession); err != nil {
		return &xhttp.ErrHttpStatus{Err: err, Status: http.StatusInternalServerError}
	}

	go a.LoginMailer.Deliver(authSession.ID.Hex(), user.Email) // Deliver later.
	return nil
}

// Login validates an auth session and
// returns an auth token in case of a successful validation.
func (a *AuthKeeper) Login(mongoSession *mgo.Session, authSessionID string,
) (*LoginResponse, error) {
	authSession, err := a.AuthSessions.FindOneByID(mongoSession, authSessionID)
	if err != nil {
		return nil, err
	}

	if err := authSession.Validate(); err != nil {
		return nil, err
	}

	authSession.LogIn()
	if err := a.AuthSessions.UpdateDoc(mongoSession, authSession); err != nil {
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
