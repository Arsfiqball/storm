package email

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

type gomailSender struct {
	dialer  *gomail.Dialer
	channel chan *gomail.Message
	config  Config
}

func NewGomailSender(d *gomail.Dialer, cfg Config) (Mailer, error) {
	if d == nil {
		return nil, fmt.Errorf("gomail dialer is nil")
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid email config: %w", err)
	}

	return &gomailSender{
		dialer:  d,
		channel: make(chan *gomail.Message, 100),
	}, nil
}

func (s *gomailSender) Start(ctx context.Context) error {
	var sc gomail.SendCloser
	var err error

	open := false

	for {
		select {
		case m, ok := <-s.channel:
			if !ok {
				return nil
			}

			if !open {
				if sc, err = s.dialer.Dial(); err != nil {
					return fmt.Errorf("failed to dial SMTP server: %w", err)
				}

				open = true
			}

			if err := gomail.Send(sc, m); err != nil {
				return fmt.Errorf("failed to send email: %w", err)
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last 30 seconds.
		case <-time.After(30 * time.Second):
			if open {
				if err := sc.Close(); err != nil {
					return fmt.Errorf("failed to close SMTP connection: %w", err)
				}

				open = false
			}
		}
	}
}

func (s *gomailSender) Stop(ctx context.Context) error {
	close(s.channel)

	return nil
}

func (s *gomailSender) Send(ctx context.Context, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.config.DefaultFrom)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	select {
	case s.channel <- m:
		return nil
	default:
		return fmt.Errorf("email channel is full")
	}
}
