package utils

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	xmiddleware "github.com/sknv/upsale/app/lib/middleware"
)

func UseDefaultMiddleware(router chi.Router) {
	router.Use(middleware.Logger, middleware.Recoverer, xmiddleware.Recoverer)
}

func RequireJWT(router chi.Router) {
	router.Use(jwtauth.Verifier(initializers.NewJWTAuth()), jwtauth.Authenticator)
}
