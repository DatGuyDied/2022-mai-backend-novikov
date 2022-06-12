package domain

import "time"

type Message struct {
	ID        int       `json:"id,omitempty"`
	From      string    `json:"from,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type PaginatedMessages struct {
	Total int        `json:"total"`
	Items []*Message `json:"items,omitempty"`
}
