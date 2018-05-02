package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/core/utils"
	"github.com/sknv/upsale/app/services"
)

type AuthSession struct {
	AuthKeeper *services.AuthKeeper
}

func NewAuthSession() *AuthSession {
	return &AuthSession{AuthKeeper: services.NewAuthKeeper()}
}

func (a *AuthSession) Create(w http.ResponseWriter, r *http.Request) {
	createRequest := a.decodeCreateRequest(w, r)
	if _, err := a.AuthKeeper.CreateAuthSession(context.Background(), createRequest); err != nil {
		log.Print("error [create auth session]: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *AuthSession) Login(w http.ResponseWriter, r *http.Request) {
	authSessionID := chi.URLParam(r, "authsessionid")
	loginResponse, err := a.AuthKeeper.Login(context.Background(), authSessionID)
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
