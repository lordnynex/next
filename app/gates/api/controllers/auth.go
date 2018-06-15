package controllers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	xchi "github.com/sknv/next/app/lib/chi"
	mongo "github.com/sknv/next/app/lib/mongo/middleware"
	xhttp "github.com/sknv/next/app/lib/net/http"
	"github.com/sknv/next/app/services"
)

const (
	authRequestLimit = 1 // Per second.
)

type Auth struct {
	AuthKeeper *services.AuthKeeper
}

func NewAuth() *Auth {
	return &Auth{AuthKeeper: services.NewAuthKeeper()}
}

func (a *Auth) Route(router chi.Router) {
	router.Route("/auth", func(r chi.Router) {
		xchi.LimitHandler(r, authRequestLimit)

		r.Post("/", a.Create)
		r.Post("/login", a.Login)
	})
}

func (a *Auth) Create(w http.ResponseWriter, r *http.Request) {
	req := a.decodeCreateRequest(w, r)
	mongoSession := mongo.GetMongoSession(r)
	if err := a.AuthKeeper.CreateAuth(mongoSession, req); err != nil {
		log.Print("[ERROR] create auth: ", err)
		err := err.(*xhttp.ErrHttpStatus)
		if err.Status != http.StatusInternalServerError {
			render.Status(r, err.Status)
			render.JSON(w, r, err.Err)
			return
		}
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	req := a.decodeLoginRequest(w, r)
	mongoSession := mongo.GetMongoSession(r)
	resp, err := a.AuthKeeper.Login(mongoSession, req)
	if err != nil {
		log.Print("[ERROR] login: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (*Auth) decodeCreateRequest(w http.ResponseWriter, r *http.Request,
) *services.CreateAuthRequest {
	req := &services.CreateAuthRequest{}
	xchi.DecodeRequest(w, r, req)
	return req
}

func (*Auth) decodeLoginRequest(w http.ResponseWriter, r *http.Request,
) *services.LoginRequest {
	req := &services.LoginRequest{}
	xchi.DecodeRequest(w, r, req)
	return req
}
