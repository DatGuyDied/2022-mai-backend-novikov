package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
)

func (m *Memory) Messages(
	ctx context.Context,
	pagination *domain.Pagination,
) (*domain.PaginatedMessages, error) {
	m.messagesMu.Lock()
	defer m.messagesMu.Unlock()

	n := len(m.messages)

	switch pagination.Order {
	case domain.PaginationOrderAsc:
		var (
			begin = (pagination.Page - 1) * pagination.Limit
			end   = pagination.Page * pagination.Limit
		)

		var page []*domain.Message
		for i, m := range m.messages {
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
				page = append(page, m.messages[i])
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

func (m *Memory) SendMessage(ctx context.Context, from string, text string) error {
	m.messagesMu.Lock()
	defer m.messagesMu.Unlock()

	m.messages = append(m.messages, &domain.Message{
		ID:        len(m.messages),
		From:      from,
		Text:      text,
		CreatedAt: time.Now().UTC(),
	})

	return nil
}
