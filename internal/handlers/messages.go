package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/auth"
)

type MessagesService interface {
	Messages(ctx context.Context, pagination *domain.Pagination) (*domain.PaginatedMessages, error)
	SendMessage(ctx context.Context, from string, text string) error
}

type Messages struct {
	service MessagesService
}

func NewMessages(service MessagesService) *Messages {
	return &Messages{
		service: service,
	}
}

func (m *Messages) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route(`/`, func(messages chi.Router) {
		messages.Get(`/`, m.messages)
		messages.Post(`/`, m.sendMessage)
	})

	return r
}

func (m *Messages) messages(w http.ResponseWriter, r *http.Request) {
	var pagination domain.Pagination
	pagination.Bind(r)

	messages, err := m.service.Messages(r.Context(), &pagination)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	render.JSON(w, r, messages)
	render.Status(r, http.StatusOK)
}

func (m *Messages) sendMessage(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("ACHTUNG: %s %s\n", user.Login, message.Text)
	if err := m.service.SendMessage(r.Context(), user.Login, message.Text); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
}
