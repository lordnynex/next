package services

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/repositories"
)

type contextKey string

const (
	contextKeyCurrentUser = contextKey("_auth.CurrentUser")
)

type (
	Authenticator struct {
		UserRepo *repositories.User
	}

	GetCurrentUserRequest struct {
		Request *http.Request
	}

	CurrentUserResponse struct {
		User *models.User
	}
)

func NewAuthenticator() *Authenticator {
	return &Authenticator{UserRepo: repositories.NewUser()}
}

func (a *Authenticator) GetCurrentUser(_ context.Context, r *GetCurrentUserRequest,
) (*CurrentUserResponse, error) {
	currentUser := r.Request.Context().Value(contextKeyCurrentUser)
	if currentUser != nil {
		currentUser := currentUser.(*models.User)
		return &CurrentUserResponse{User: currentUser}, nil
	}

	_, claims, _ := jwtauth.FromContext(r.Request.Context())
	userID := claims["sub"].(string)
	user, err := a.UserRepo.FindOneByID(userID)
	if err != nil {
		return nil, err
	}

	// Cache current user.
	*r.Request = *r.Request.WithContext(
		context.WithValue(r.Request.Context(), contextKeyCurrentUser, user),
	)
	return &CurrentUserResponse{User: user}, nil
}
