package main

import (
	"time"

	"github.com/go-chi/chi"

	"github.com/sknv/upsale/app/core/utils"
	"github.com/sknv/upsale/app/gates/api/cfg"
	"github.com/sknv/upsale/app/gates/api/controllers"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
)

const (
	concurrentRequestLimit = 1000
	requestTimeout         = 60 * time.Second
	shutdownTimeout        = 30 * time.Second
)

func main() {
	router := chi.NewRouter()
	utils.UseDefaultMiddleware(router)
	utils.ThrottleAndTimeout(router, concurrentRequestLimit, requestTimeout)

	route(router)
	xhttp.ListenAndServe(cfg.GetAddr(), router, shutdownTimeout)
}

func route(router chi.Router) {
	controllers.NewAuthSession().Route(router)
	controllers.NewGreeter().Route(router)
}
