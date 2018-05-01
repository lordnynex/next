package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sknv/upsale/app/core/cfg"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
	"github.com/sknv/upsale/app/services/session"
)

type Session struct {
	SessionClient session.Session
}

func NewSession() *Session {
	return &Session{
		SessionClient: session.NewSessionClient(cfg.GetSecretKey()),
	}
}

func (s *Session) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := s.decodeLoginRequest(w, r)

	loginResponse, err := s.SessionClient.Login(context.Background(), loginRequest)
	if err != nil {
		log.Print("error [login]: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		xhttp.AbortHandler()
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, loginResponse)
}

func (*Session) decodeLoginRequest(w http.ResponseWriter, r *http.Request) *session.LoginRequest {
	loginRequest := &session.LoginRequest{}
	err := render.DecodeJSON(r.Body, loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		xhttp.AbortHandler()
	}
	return loginRequest
}
