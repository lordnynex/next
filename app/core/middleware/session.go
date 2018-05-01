package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/sknv/upsale/app/services/session"
)

func SessionVerifier(next http.Handler) http.Handler {
	sessionClient := session.NewSessionClient()
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, err := sessionClient.GetUserID(
			context.Background(), &session.GetUserIDRequest{Context: r.Context()},
		)
		if err != nil {
			log.Print("error [verify session]: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
