package services

import (
	"net/http"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-chi/jwtauth"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/sknv/next/app/core/initers"
	"github.com/sknv/next/app/core/mailers"
	"github.com/sknv/next/app/core/store"
	xhttp "github.com/sknv/next/app/lib/net/http"
)

const (
	exp = 90 * 24 * time.Hour // Auth session expires in 90 days.
)

type (
	AuthKeeper struct {
		JWTAuth     *jwtauth.JWTAuth
		LoginMailer *mailers.Login
		Users       *store.User
	}

	CreateAuthRequest struct {
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	LoginResponse struct {
		AuthToken string `json:"authtoken"`
	}
)

func NewAuthKeeper() *AuthKeeper {
	return &AuthKeeper{
		JWTAuth:     initers.GetJWTAuth(),
		LoginMailer: mailers.NewLogin(),
		Users:       store.NewUser(),
	}
}

// CreateAuth creates a user account if one does not exist yet and stores an auth session.
func (a *AuthKeeper) CreateAuth(mongoSession *mgo.Session, r *CreateAuthRequest) error {
	if err := r.Validate(); err != nil {
		return &xhttp.ErrHttpStatus{Err: err, Status: http.StatusUnprocessableEntity}
	}

	// Create an account if one does not exist yet.
	user, err := a.Users.FindOneOrInsertByEmail(mongoSession, r.Email)
	if err != nil {
		return &xhttp.ErrHttpStatus{Err: err, Status: http.StatusInternalServerError}
	}

	user.GenerateCode()
	if err := a.Users.UpdateDoc(mongoSession, user); err != nil {
		return &xhttp.ErrHttpStatus{Err: err, Status: http.StatusInternalServerError}
	}

	go a.LoginMailer.Deliver(user) // Deliver later.
	return nil
}

// Login authenticates an auth session.
func (a *AuthKeeper) Login(mongoSession *mgo.Session, r *LoginRequest,
) (*LoginResponse, error) {
	user, err := a.Users.FindOneByEmail(mongoSession, r.Email)
	if err != nil {
		return nil, err
	}

	if err := user.Authenticate(r.Code); err != nil {
		return nil, err
	}

	user.LogIn()
	if err := a.Users.UpdateDoc(mongoSession, user); err != nil {
		return nil, err
	}

	_, tokenString, err := a.JWTAuth.Encode(
		jwtauth.Claims{"sub": user.ID.Hex(), "exp": time.Now().Add(exp).Unix()},
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{AuthToken: tokenString}, nil
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

func (r *CreateAuthRequest) Validate() error {
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Email, validation.Required, is.Email),
	)
}
