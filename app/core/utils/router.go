package utils

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/initializers"
	coremiddleware "github.com/sknv/upsale/app/core/middleware"
	libmiddleware "github.com/sknv/upsale/app/lib/middleware"
)

func UseDefaultMiddleware(router chi.Router) {
	router.Use(
		middleware.RealIP, middleware.Logger, middleware.Recoverer, libmiddleware.Recoverer,
	)
}

func ThrottleAndTimeout(
	router chi.Router, concurrentRequestLimit int, requestTimeout time.Duration,
) {
	router.Use(
		middleware.Throttle(concurrentRequestLimit), middleware.Timeout(requestTimeout),
	)
}

func LimitHandler(router chi.Router, requestLimit float64) {
	router.Use(tollbooth_chi.LimitHandler(tollbooth.NewLimiter(requestLimit, nil)))
}

func RequireLogin(router chi.Router) {
	router.Use(
		// Require presence of valid JWT.
		jwtauth.Verifier(initializers.GetJWTAuth()), jwtauth.Authenticator,
		// Require presence of current user.
		coremiddleware.CurrentUserVerifier,
	)
}
