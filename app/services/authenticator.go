package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/store"
)

type contextKey string

const (
	contextKeyCurrentUser = contextKey("authenticator.currentuser")
)

type Authenticator struct {
	Users *store.User
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{Users: store.NewUser()}
}

func (a *Authenticator) GetCurrentUser(_ context.Context, r *http.Request,
) (*models.User, error) {
	currentUser := r.Context().Value(contextKeyCurrentUser)
	if currentUser != nil {
		currentUser := currentUser.(*models.User)
		return currentUser, nil
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("sub claim is empty or not a string")
	}

	user, err := a.Users.FindOneByID(nil, userID)
	if err != nil {
		return nil, err
	}

	// Cache current user.
	*r = *r.WithContext(context.WithValue(r.Context(), contextKeyCurrentUser, user))
	return user, nil
}
