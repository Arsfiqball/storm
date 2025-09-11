package provider

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

type Gomail interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type gomailState struct {
	dialer  *gomail.Dialer
	channel chan *gomail.Message
}

func (g *gomailState) Start(ctx context.Context) error {
	var sc gomail.SendCloser
	var err error

	open := false

	for {
		select {
		case m, ok := <-g.channel:
			if !ok {
				return nil
			}

			if !open {
				if sc, err = g.dialer.Dial(); err != nil {
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

func (g *gomailState) Stop(ctx context.Context) error {
	close(g.channel)

	return nil
}

func ProvideGomail() (Gomail, error) {
	dialer := gomail.NewDialer("smtp.example.com", 587, "user", "password")

	return &gomailState{
		dialer:  dialer,
		channel: make(chan *gomail.Message, 100),
	}, nil
}
