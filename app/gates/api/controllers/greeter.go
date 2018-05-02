package controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/services"
)

type Greeter struct {
	Authenticator *services.Authenticator
}

func NewGreeter() *Greeter {
	return &Greeter{Authenticator: services.NewAuthenticator()}
}

func (g *Greeter) Hello(w http.ResponseWriter, r *http.Request) {
	currentUserResponse, _ := g.Authenticator.GetCurrentUser(
		context.Background(), &services.GetCurrentUserRequest{Request: r},
	)
	render.JSON(w, r, render.M{"hello": currentUserResponse.User.Username})
}
