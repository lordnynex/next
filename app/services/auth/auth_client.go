package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	"github.com/sknv/upsale/app/core/models"
	"github.com/sknv/upsale/app/core/repositories"
)

type contextKey string

const (
	contextKeyCurrentUser = contextKey("_auth.CurrentUser")
	exp                   = 90 * 24 * time.Hour // Expires in 90 days.
)

type AuthClient struct {
	JWTAuth  *jwtauth.JWTAuth
	UserRepo *repositories.User
}

func NewAuthClient() Auth {
	return &AuthClient{
		JWTAuth:  initializers.NewJWTAuth(),
		UserRepo: repositories.NewUser(),
	}
}

func (c *AuthClient) Login(_ context.Context, r *LoginRequest) (*LoginResponse, error) {
	r.Login = strings.TrimSpace(r.Login)
	if r.Login != "login" {
		return nil, errors.New("user is not authenticated")
	}

	userID := "qwe123"
	_, tokenString, err := c.JWTAuth.Encode(
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

func (c *AuthClient) GetCurrentUser(_ context.Context, r *GetCurrentUserRequest,
) (*CurrentUserResponse, error) {
	currentUser := r.Request.Context().Value(contextKeyCurrentUser)
	if currentUser != nil {
		currentUser := currentUser.(*models.User)
		return &CurrentUserResponse{User: currentUser}, nil
	}

	_, claims, _ := jwtauth.FromContext(r.Request.Context())
	userID := claims["sub"].(string)
	user, err := c.UserRepo.FindOneByID(userID)
	if err != nil {
		return nil, err
	}

	// Cache current user.
	*r.Request = *r.Request.WithContext(
		context.WithValue(r.Request.Context(), contextKeyCurrentUser, user),
	)
	return &CurrentUserResponse{User: user}, nil
}
