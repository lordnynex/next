package controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/services/auth"
)

type Greeter struct {
	AuthClient auth.Auth
}

func NewGreeter() *Greeter {
	return &Greeter{AuthClient: auth.NewAuthClient()}
}

func (g *Greeter) Hello(w http.ResponseWriter, r *http.Request) {
	currentUserResponse, _ := g.AuthClient.GetCurrentUser(
		context.Background(), &auth.GetCurrentUserRequest{Context: r.Context()},
	)
	render.JSON(w, r, render.M{"hello": currentUserResponse.User.Username})
}
