package main

import (
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sknv/upsale/app/gates/api/cfg"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
)

const (
	shutdownTimeout = 10 * time.Second
)

func main() {
	addr := cfg.GetAddr()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/hello", hello)

	xhttp.ListenAndServe(addr, router, shutdownTimeout)
}

func hello(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{"name": "User"})
}
