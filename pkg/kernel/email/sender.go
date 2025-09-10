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

type Mailer interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Send(ctx context.Context, to, subject, body string) error
}
