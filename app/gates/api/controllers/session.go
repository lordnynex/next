package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"

	xhttp "github.com/sknv/upsale/app/lib/net/http"
	"github.com/sknv/upsale/app/services"
)

type Session struct {
	Keeper *services.Keeper
}

func NewSession() *Session {
	return &Session{Keeper: services.NewKeeper()}
}

func (s *Session) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := s.decodeLoginRequest(w, r)

	loginResponse, err := s.Keeper.Login(context.Background(), loginRequest)
	if err != nil {
		log.Print("error [login]: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		xhttp.AbortHandler()
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, loginResponse)
}

func (*Session) decodeLoginRequest(w http.ResponseWriter, r *http.Request,
) *services.LoginRequest {
	loginRequest := &services.LoginRequest{}
	err := render.DecodeJSON(r.Body, loginRequest)
	if err != nil {
		log.Print("error [decodeLoginRequest]: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		xhttp.AbortHandler()
	}
	return loginRequest
}
