package services

import (
	"context"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
)

type UserMessagesRepository interface {
	UserMessages(ctx context.Context, login string, pagination *domain.Pagination) (*domain.PaginatedMessages, error)
	SendUserMessage(ctx context.Context, login string, to string, text string) error
}

type UserMessagesService struct {
	repo UserMessagesRepository
}

func NewUserMessagesService(repo UserMessagesRepository) *UserMessagesService {
	return &UserMessagesService{
		repo: repo,
	}
}

func (s *UserMessagesService) UserMessages(
	ctx context.Context,
	login string,
	pagination *domain.Pagination,
) (*domain.PaginatedMessages, error) {
	return s.repo.UserMessages(ctx, login, pagination)
}

func (s *UserMessagesService) SendUserMessage(
	ctx context.Context,
	login, to, text string,
) error {
	return s.repo.SendUserMessage(ctx, login, to, text)
}
