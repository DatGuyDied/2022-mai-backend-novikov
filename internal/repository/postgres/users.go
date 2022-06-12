package postgres

import (
	"context"
	"fmt"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/user"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/usermessage"
)

func (p *postgres) CreateUser(ctx context.Context, u *domain.User) error {
	users, err := p.client.User.
		Query().
		Where(user.Login(u.Login)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}
	if len(users) != 0 {
		return domain.ErrAlreadyExists
	}

	_, err = p.client.User.
		Create().
		SetLogin(u.Login).
		SetHash(u.Hash).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("can not create user: %w", err)
	}

	return nil
}

func (p *postgres) UserHash(ctx context.Context, login string) ([]byte, error) {
	users, err := p.client.User.
		Query().
		Where(user.Login(login)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not get user with login `%s`: %w", login, err)
	}
	if len(users) == 0 {
		return nil, domain.ErrNotExists
	}

	return users[0].Hash, nil
}

func (p *postgres) UserMessages(
	ctx context.Context,
	login string,
	pagination *domain.Pagination,
) (*domain.PaginatedMessages, error) {
	entMessages, err := p.client.UserMessage.
		Query().
		Where(usermessage.To(login)).
		Limit(pagination.Limit).
		Offset(pagination.Limit * (pagination.Page - 1)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user messages: %w", err)
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

func (p *postgres) SendUserMessage(ctx context.Context, login string, to string, text string) error {
	_, err := p.client.UserMessage.
		Create().
		SetFrom(login).
		SetTo(to).
		SetText(text).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("create user message: %w", err)
	}

	return nil
}
