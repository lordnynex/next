package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/sknv/next/app/core/utils"
	"github.com/sknv/next/app/services"
)

type User struct {
	UserService *services.User
}

func NewUser() *User {
	return &User{UserService: services.NewUser()}
}

func (u *User) Route(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		utils.RequireLogin(r)

		r.Get("/me", u.Me)
	})
}

func (u *User) Me(w http.ResponseWriter, r *http.Request) {
	resp, _ := u.UserService.Me(r)
	render.JSON(w, r, resp)
}
