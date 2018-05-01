package controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/services/session"
)

type Greeter struct {
	SessionClient session.Session
}

func NewGreeter() *Greeter {
	return &Greeter{SessionClient: session.NewSessionClient()}
}

func (g *Greeter) Hello(w http.ResponseWriter, r *http.Request) {
	userIDResponse, _ := g.SessionClient.GetUserID(
		context.Background(), &session.GetUserIDRequest{Context: r.Context()},
	)
	render.JSON(w, r, render.M{"hello": userIDResponse.UserID})
}
