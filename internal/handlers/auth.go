package handlers

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AuthService interface {
	Login(ctx context.Context, uc *domain.UserCredentials) error
	Register(ctx context.Context, uc *domain.UserCredentials) error
}

type Auth struct {
	service    AuthService
	privateKey *rsa.PrivateKey
}

func NewAuth(service AuthService, pk *rsa.PrivateKey) *Auth {
	return &Auth{
		service:    service,
		privateKey: pk,
	}
}

func (a *Auth) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route(`/`, func(auth chi.Router) {
		auth.Post(`/login`, a.login)
		auth.Post(`/register`, a.register)
	})

	return r
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	var uc domain.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	loginErr := a.service.Login(r.Context(), &uc)
	if loginErr != nil {
		render.Status(r, http.StatusUnauthorized)
		render.PlainText(w, r, loginErr.Error())
		return
	}

	user := &auth.User{
		Login: uc.Login,
	}
	if err := auth.WriteToken(w, user, a.privateKey); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
}

func (a *Auth) register(w http.ResponseWriter, r *http.Request) {
	var uc domain.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	if err := a.service.Register(r.Context(), &uc); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	fmt.Println(uc)

	render.Status(r, http.StatusOK)
}
