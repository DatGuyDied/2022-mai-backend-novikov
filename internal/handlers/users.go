package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/auth"
)

type UserMessagesService interface {
	UserMessages(ctx context.Context, login string, pagination *domain.Pagination) (*domain.PaginatedMessages, error)
	SendUserMessage(ctx context.Context, from string, to string, text string) error
}

type Users struct {
	service UserMessagesService
}

func NewUsers(userMessages UserMessagesService) *Users {
	return &Users{
		service: userMessages,
	}
}

func (u *Users) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route(`/`, func(users chi.Router) {
		users.Route(`/me`, func(me chi.Router) {
			me.Get(`/messages`, u.messages)
		})
		users.Route(`/{id}`, func(otherUser chi.Router) {
			otherUser.Post(`/messages`, u.sendMessage)
		})
	})

	return r
}

func (u *Users) messages(w http.ResponseWriter, r *http.Request) {
	var pagination domain.Pagination
	pagination.Bind(r)

	user := auth.UserFromContext(r.Context())
	messages, err := u.service.UserMessages(r.Context(), user.Login, &pagination)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	render.JSON(w, r, messages)
	render.Status(r, http.StatusOK)
}

func (u *Users) sendMessage(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Text string `json:"text,omitempty"`
	}
	var message Message
	if err := render.Decode(r, &message); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	user := auth.UserFromContext(r.Context())
	to := chi.URLParam(r, `id`)
	sendErr := u.service.SendUserMessage(
		r.Context(),
		user.Login,
		to,
		message.Text,
	)
	if sendErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, sendErr.Error())
		return
	}

	render.Status(r, http.StatusOK)
}
