package middleware

import (
	"log"
	"net/http"

	"github.com/sknv/upsale/app/services"
)

func CurrentUserVerifier(next http.Handler) http.Handler {
	authenticator := services.NewAuthenticator()
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, err := authenticator.GetCurrentUser(r)
		if err != nil {
			log.Print("[ERROR] verify current user: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
