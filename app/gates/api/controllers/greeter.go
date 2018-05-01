package controllers

import (
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

type Greeter struct{}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (*Greeter) Hello(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	render.JSON(w, r, render.M{"hello": claims["sub"]})
}
