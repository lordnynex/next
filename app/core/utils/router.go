package utils

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	xmiddleware "github.com/sknv/upsale/app/core/middleware"
)

func RequireLogin(router chi.Router) {
	router.Use(
		// Require presence of valid JWT.
		jwtauth.Verifier(initializers.GetJWTAuth()), jwtauth.Authenticator,
		// Require presence of current user.
		xmiddleware.CurrentUserVerifier,
	)
}
