package memory

import (
	"sync"

	"github.com/DatGuyDied/2022-mai-backend-novikov/domain"
)

type Memory struct {
	usersMu sync.Mutex
	users   map[string]*domain.User

	messagesMu sync.Mutex
	messages   []*domain.Message

	userMessagesMu sync.Mutex
	userMessages   map[string][]*domain.Message
}

func New() *Memory {
	return &Memory{
		users:        make(map[string]*domain.User),
		userMessages: make(map[string][]*domain.Message),
	}
}
