package memory

import (
	"context"
	"fmt"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
)

func (m *Memory) CreateUser(ctx context.Context, u *domain.User) error {
	m.usersMu.Lock()
	defer m.usersMu.Unlock()

	if _, ok := m.users[u.Login]; ok {
		return domain.ErrAlreadyExists
	}

	m.users[u.Login] = u
	return nil
}

func (m *Memory) UserHash(ctx context.Context, login string) ([]byte, error) {
	m.usersMu.Lock()
	defer m.usersMu.Unlock()

	u, ok := m.users[login]
	if !ok {
		return nil, domain.ErrNotExists
	}

	return u.Hash, nil
}

func (m *Memory) UserMessages(
	ctx context.Context,
	login string,
	pagination *domain.Pagination,
) (*domain.PaginatedMessages, error) {
	m.userMessagesMu.Lock()
	defer m.userMessagesMu.Unlock()

	userMessages := m.userMessages[login]
	n := len(userMessages)

	switch pagination.Order {
	case domain.PaginationOrderAsc:
		var (
			begin = (pagination.Page - 1) * pagination.Limit
			end   = pagination.Page * pagination.Limit
		)

		var page []*domain.Message
		for i, m := range userMessages {
			if i >= begin && i < end {
				page = append(page, m)
			}
		}

		return &domain.PaginatedMessages{
			Total: len(page),
			Items: page,
		}, nil

	case domain.PaginationOrderDesc:
		var (
			begin = n - 1 - (pagination.Page-1)*pagination.Limit
			end   = n - 1 - pagination.Page*pagination.Limit
		)

		var page []*domain.Message
		for i := n - 1; i >= 0; i-- {
			if i <= begin && i > end {
				page = append(page, userMessages[i])
			}
		}

		return &domain.PaginatedMessages{
			Total: len(page),
			Items: page,
		}, nil

	default:
		return nil, fmt.Errorf("%w: %v",
			domain.ErrInvalidPaginationParams,
			"invalid order",
		)
	}
}

func (m *Memory) SendUserMessage(
	ctx context.Context,
	login, to, text string,
) error {
	m.userMessagesMu.Lock()
	defer m.userMessagesMu.Unlock()

	if _, ok := m.users[to]; !ok {
		return domain.ErrNotExists
	}

	id := len(m.userMessages[to])
	message := &domain.Message{
		ID:   id,
		From: login,
		Text: text,
	}

	m.userMessages[login] = append(m.userMessages[login], message)
	m.userMessages[to] = m.userMessages[login]

	return nil
}
