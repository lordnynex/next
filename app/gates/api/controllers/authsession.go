package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/core/utils"
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
	if _, err := a.Keeper.CreateAuthSession(context.Background(), createRequest); err != nil {
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
	utils.DecodeRequest(w, r, createRequest)
	return createRequest
}

func (*AuthSession) decodeLoginRequest(w http.ResponseWriter, r *http.Request,
) *services.LoginRequest {
	loginRequest := &services.LoginRequest{}
	utils.DecodeRequest(w, r, loginRequest)
	return loginRequest
}
