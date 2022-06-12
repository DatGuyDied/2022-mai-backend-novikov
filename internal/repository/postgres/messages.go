package postgres

import (
	"context"
	"fmt"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent"
)

func (p *postgres) Messages(ctx context.Context, pagination *domain.Pagination) (*domain.PaginatedMessages, error) {
	entMessages, err := p.client.Message.
		Query().
		Limit(pagination.Limit).
		Offset(pagination.Limit * (pagination.Page - 1)).
		Order(ent.Asc("id")).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("get messages: %w", err)
	}

	messages := make([]*domain.Message, 0, len(entMessages))
	for _, m := range entMessages {
		messages = append(messages, &domain.Message{
			ID:        m.ID,
			From:      m.From,
			Text:      m.Text,
			CreatedAt: m.CtreatedAt,
		})
	}

	return &domain.PaginatedMessages{
		Total: len(messages),
		Items: messages,
	}, nil
}

func (p *postgres) SendMessage(ctx context.Context, from string, text string) error {
	_, err := p.client.Message.
		Create().
		SetFrom(from).
		SetText(text).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("create message: %w", err)
	}

	return nil
}
