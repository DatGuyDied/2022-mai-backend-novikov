package repository

import (
	"context"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/repository/memory"
	"github.com/DatGuyDied/2022-mai-backend-novikov/internal/repository/postgres"
)

func NewMemoryRepository() Repository {
	return memory.New()
}

func NewPostgresRepository(ctx context.Context) (Repository, error) {
	return postgres.New(ctx)
}

type Repository interface {
	UserHash(ctx context.Context, login string) ([]byte, error)
	CreateUser(ctx context.Context, u *domain.User) error

	Messages(ctx context.Context, pagination *domain.Pagination) (*domain.PaginatedMessages, error)
	SendMessage(ctx context.Context, from string, text string) error

	UserMessages(ctx context.Context, login string, pagination *domain.Pagination) (*domain.PaginatedMessages, error)
	SendUserMessage(ctx context.Context, login string, to string, text string) error
}
