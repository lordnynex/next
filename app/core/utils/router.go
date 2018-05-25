package utils

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	cmiddleware "github.com/sknv/upsale/app/core/middleware"
	lmiddleware "github.com/sknv/upsale/app/lib/middleware"
)

func UseDefaultMiddleware(router chi.Router) {
	router.Use(
		middleware.RealIP, middleware.Logger, middleware.Recoverer, lmiddleware.Recoverer,
	)
}

func UseThrottleAndTimeout(
	router chi.Router, concurrentRequestsLimit int, requestTimeout time.Duration,
) {
	router.Use(
		middleware.Throttle(concurrentRequestsLimit), middleware.Timeout(requestTimeout),
	)
}

func RequireLogin(router chi.Router) {
	router.Use(
		// Require presence of valid JWT.
		jwtauth.Verifier(initializers.GetJWTAuth()), jwtauth.Authenticator,
		// Require presence of current user.
		cmiddleware.CurrentUserVerifier,
	)
}
