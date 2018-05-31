package middleware

import (
	"context"
	"net/http"

	"github.com/globalsign/mgo"
)

type contextKey string

const (
	contextKeyMongoSession = contextKey("mongosession")
)

// WithMongoSession puts a Mongo session instance to a request context.
func WithMongoSession(session *mgo.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Copy a Mongo session and schedule a clean up.
			sessionCopy := session.Copy()
			defer sessionCopy.Close()

			// Put the session into a request context.
			ctx := context.WithValue(r.Context(), contextKeyMongoSession, sessionCopy)

			// Process request.
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// GetMongoSession returns a Mongo session from a request context.
func GetMongoSession(r *http.Request) *mgo.Session {
	return r.Context().Value(contextKeyMongoSession).(*mgo.Session)
}
