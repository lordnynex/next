package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"

	xhttp "github.com/sknv/upsale/app/lib/net/http"
	"github.com/sknv/upsale/app/services"
)

type AuthSession struct {
	Keeper *services.Keeper
}

func NewAuthSession() *AuthSession {
	return &AuthSession{Keeper: services.NewKeeper()}
}

func (a *AuthSession) Create(w http.ResponseWriter, r *http.Request) {
	createRequest := a.decodeCreateRequest(w, r)
	_, err := a.Keeper.CreateAuthSession(context.Background(), createRequest)
	if err != nil {
		if err != services.ErrUserDoesNotExist {
			panic(err)
		}

		log.Print("error [create auth session]: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *AuthSession) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := a.decodeLoginRequest(w, r)
	loginResponse, err := a.Keeper.Login(context.Background(), loginRequest)
	if err != nil {
		log.Print("error [login]: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, loginResponse)
}

func (*AuthSession) decodeCreateRequest(w http.ResponseWriter, r *http.Request,
) *services.CreateAuthSessionRequest {
	createRequest := &services.CreateAuthSessionRequest{}
	err := render.DecodeJSON(r.Body, createRequest)
	if err != nil {
		log.Print("error [decodeCreateRequest]: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		xhttp.AbortHandler()
	}
	return createRequest
}

func (*AuthSession) decodeLoginRequest(w http.ResponseWriter, r *http.Request,
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
