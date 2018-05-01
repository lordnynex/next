package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/sknv/upsale/app/services/auth"
)

func CurrentUserVerifier(next http.Handler) http.Handler {
	authClient := auth.NewAuthClient()
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, err := authClient.GetCurrentUser(
			context.Background(), &auth.GetCurrentUserRequest{Request: r},
		)
		if err != nil {
			log.Print("error [verify current user]: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
