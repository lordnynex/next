package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"

	xhttp "github.com/sknv/upsale/app/lib/net/http"
	"github.com/sknv/upsale/app/services/auth"
)

type Auth struct {
	AuthClient auth.Auth
}

func NewAuth() *Auth {
	return &Auth{AuthClient: auth.NewAuthClient()}
}

func (s *Auth) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := s.decodeLoginRequest(w, r)

	loginResponse, err := s.AuthClient.Login(context.Background(), loginRequest)
	if err != nil {
		log.Print("error [login]: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		xhttp.AbortHandler()
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, loginResponse)
}

func (*Auth) decodeLoginRequest(w http.ResponseWriter, r *http.Request) *auth.LoginRequest {
	loginRequest := &auth.LoginRequest{}
	err := render.DecodeJSON(r.Body, loginRequest)
	if err != nil {
		log.Print("error [decodeLoginRequest]: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		xhttp.AbortHandler()
	}
	return loginRequest
}
