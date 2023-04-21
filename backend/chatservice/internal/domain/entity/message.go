package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/j178/tiktoken-go"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role string, content string, model *Model) (*Message, error) {
	totalTokens := tiktoken_go.CountTokens(model.GetModelName(), content)
	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Tokens:    totalTokens,
		Model:     model,
		CreatedAt: time.Now(),
	}
	if err := msg.Validate(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (msg *Message) Validate() error {
	if msg.Role != "user" && msg.Role != "system" && msg.Role != "assistant" {
		return errors.New("invalid role")
	}

	if msg.Content == "" {
		return errors.New("content is empty")
	}

	if msg.CreatedAt.IsZero() {
		return errors.New("created at invalid")
	}
	return nil
}

func (msg *Message) GetQtdTokens() int {
	return msg.Tokens
}
