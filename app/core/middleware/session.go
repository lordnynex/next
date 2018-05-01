package middleware

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"

	"github.com/sknv/upsale/app/core/repositories"
	"github.com/sknv/upsale/app/services/session"
)

func SessionVerifier(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		sessionID, ok := claims[session.ClaimSessionID].(string)
		if !ok {
			log.Print("error [verify session]: session id does not exist in jwt claims")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		sessionRepo := &repositories.Session{}
		_, err := sessionRepo.FindOneByID(sessionID)
		if err != nil {
			log.Print("error [verify session]: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
