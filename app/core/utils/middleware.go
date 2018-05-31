package utils

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/next/app/core/initers"
	core "github.com/sknv/next/app/core/middleware"
)

func RequireLogin(router chi.Router) {
	router.Use(
		// Require presence of valid JWT.
		jwtauth.Verifier(initers.GetJWTAuth()), jwtauth.Authenticator,
		// Require presence of current user.
		core.CurrentUserVerifier,
	)
}
