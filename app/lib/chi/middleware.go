package chi

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_chi"
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	lib "github.com/sknv/next/app/lib/middleware"
	mongo "github.com/sknv/next/app/lib/mongo/middleware"
)

func UseDefaultMiddleware(router chi.Router) {
	router.Use(
		middleware.RealIP, middleware.Logger, middleware.Recoverer, lib.Recoverer,
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

func ProvideMongoSession(router chi.Router, session *mgo.Session) {
	router.Use(mongo.WithMongoSession(session))
}
