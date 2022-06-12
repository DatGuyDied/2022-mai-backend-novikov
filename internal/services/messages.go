package services

import (
	"context"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
)

type MessagesRepository interface {
	Messages(ctx context.Context, pagination *domain.Pagination) (*domain.PaginatedMessages, error)
	SendMessage(ctx context.Context, from string, text string) error
}

type MessagesService struct {
	repo MessagesRepository
}

func NewMessagesService(repo MessagesRepository) *MessagesService {
	return &MessagesService{
		repo: repo,
	}
}

func (s *MessagesService) Messages(
	ctx context.Context,
	pagination *domain.Pagination,
) (*domain.PaginatedMessages, error) {
	return s.repo.Messages(ctx, pagination)
}

func (s *MessagesService) SendMessage(
	ctx context.Context,
	from string,
	text string,
) error {
	return s.repo.SendMessage(ctx, from, text)
}
