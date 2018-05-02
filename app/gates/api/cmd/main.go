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
	shutdownTimeout = 10 * time.Second
)

func main() {
	addr := cfg.GetAddr()

	router := chi.NewRouter()
	utils.UseDefaultMiddleware(router)

	route(router)
	xhttp.ListenAndServe(addr, router, shutdownTimeout)
}

func route(router chi.Router) {
	routeAuth(router)
	routeGreeter(router)
}

func routeAuth(router chi.Router) {
	session := controllers.NewSession()
	router.Post("/login", session.Login)
}

func routeGreeter(router chi.Router) {
	greeter := controllers.NewGreeter()
	router.Route("/greeter", func(r chi.Router) {
		utils.RequireLogin(r)

		r.Get("/hello", greeter.Hello)
	})
}
