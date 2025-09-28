package mailer

import "context"

type DefaultFrom struct {
	Name  string
	Email string
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
