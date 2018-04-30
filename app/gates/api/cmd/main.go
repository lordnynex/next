package main

import (
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sknv/upsale/app/gates/api/cfg"
	xmiddleware "github.com/sknv/upsale/app/lib/middleware"
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
	router.Use(xmiddleware.Recoverer)

	router.Get("/hello", hello)
	router.Get("/abort", abort)

	xhttp.ListenAndServe(addr, router, shutdownTimeout)
}

func hello(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{"name": "User"})
}

func abort(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, render.M{"error": "Unauthorized"})
	xhttp.AbortHandler()
}
