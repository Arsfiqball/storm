package email

import (
	"context"
	"errors"
)

type Config struct {
	DefaultFrom string
}

func (c Config) Validate() error {
	if c.DefaultFrom == "" {
		return errors.New("default from email is required")
	}

	return nil
}

type Message struct {
	From    string // Optional. If empty, use Config.DefaultFrom
	To      string
	Subject string
	Body    string
}

type Mailer interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Send(ctx context.Context, m Message) error
}
