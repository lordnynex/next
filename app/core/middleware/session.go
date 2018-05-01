package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/sknv/upsale/app/lib/net/rpc"
	"github.com/sknv/upsale/app/services/session"
)

func SessionVerifier(next http.Handler) http.Handler {
	sessionClient := session.NewSessionClient()
	fn := func(w http.ResponseWriter, r *http.Request) {
		idResponse, err := sessionClient.IDFromContext(r.Context(), &rpc.Empty{})
		if err != nil {
			log.Print("error [verify session]: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		_, err = sessionClient.FindOneByID(
			context.Background(), &session.FindOneByIDRequest{ID: idResponse.SessionID},
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
