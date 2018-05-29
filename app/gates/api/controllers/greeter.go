package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/core/utils"
	"github.com/sknv/upsale/app/services"
)

type Greeter struct {
	Authenticator *services.Authenticator
}

func NewGreeter() *Greeter {
	return &Greeter{Authenticator: services.NewAuthenticator()}
}

func (g *Greeter) Route(router chi.Router) {
	router.Route("/greeter", func(r chi.Router) {
		utils.RequireLogin(r)

		r.Get("/hello", g.Hello)
	})
}

func (g *Greeter) Hello(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := g.Authenticator.GetCurrentUser(r)
	render.JSON(w, r, render.M{"hello": currentUser.Email})
}
