package utils

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	cmiddleware "github.com/sknv/upsale/app/core/middleware"
	lmiddleware "github.com/sknv/upsale/app/lib/middleware"
)

func UseDefaultMiddleware(router chi.Router) {
	router.Use(middleware.Logger, middleware.Recoverer, lmiddleware.Recoverer)
}

func RequireJWT(router chi.Router) {
	router.Use(jwtauth.Verifier(initializers.NewJWTAuth()), jwtauth.Authenticator)
}

func RequireLogin(router chi.Router) {
	RequireJWT(router)
	router.Use(cmiddleware.CurrentUserVerifier)
}
