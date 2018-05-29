package controllers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/sknv/upsale/app/core/utils"
	xchi "github.com/sknv/upsale/app/lib/chi"
	mongo "github.com/sknv/upsale/app/lib/mongo/middleware"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
	"github.com/sknv/upsale/app/services"
)

const (
	authRequestLimit = 1 // Per second.
)

type AuthSession struct {
	AuthKeeper *services.AuthKeeper
}

func NewAuthSession() *AuthSession {
	return &AuthSession{AuthKeeper: services.NewAuthKeeper()}
}

func (a *AuthSession) Route(router chi.Router) {
	router.Route("/login", func(r chi.Router) {
		xchi.LimitHandler(r, authRequestLimit)

		r.Post("/", a.Create)
		r.Post("/{authsessionid}", a.Login)
	})
}

func (a *AuthSession) Create(w http.ResponseWriter, r *http.Request) {
	req := a.decodeCreateRequest(w, r)
	mongoSession := mongo.GetMongoSession(r)
	if err := a.AuthKeeper.CreateAuthSession(mongoSession, req); err != nil {
		log.Print("[ERROR] create auth session: ", err)
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

func (a *AuthSession) Login(w http.ResponseWriter, r *http.Request) {
	authSessionID := chi.URLParam(r, "authsessionid")
	mongoSession := mongo.GetMongoSession(r)
	resp, err := a.AuthKeeper.Login(mongoSession, authSessionID)
	if err != nil {
		log.Print("[ERROR] login: ", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (*AuthSession) decodeCreateRequest(w http.ResponseWriter, r *http.Request,
) *services.CreateAuthSessionRequest {
	req := &services.CreateAuthSessionRequest{}
	utils.DecodeRequest(w, r, req)
	return req
}
