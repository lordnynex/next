package main

import (
	"time"

	"github.com/go-chi/chi"

	"github.com/sknv/upsale/app/core/initers"
	"github.com/sknv/upsale/app/gates/api/cfg"
	"github.com/sknv/upsale/app/gates/api/controllers"
	xchi "github.com/sknv/upsale/app/lib/chi"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
)

const (
	concurrentRequestLimit = 1000
	requestTimeout         = 60 * time.Second
	shutdownTimeout        = 30 * time.Second
)

func main() {
	mongoSession := initers.GetMongoSession()
	defer mongoSession.Close() // Clean up.

	router := chi.NewRouter()
	xchi.UseDefaultMiddleware(router)
	xchi.ThrottleAndTimeout(router, concurrentRequestLimit, requestTimeout)
	xchi.ProvideMongoSession(router, mongoSession)

	route(router)
	xhttp.ListenAndServe(cfg.GetAddr(), router, shutdownTimeout)
}

func route(router chi.Router) {
	controllers.NewAuthSession().Route(router)
	controllers.NewGreeter().Route(router)
}
