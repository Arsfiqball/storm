package provider

import (
	"app/pkg/kernel/email"
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type Email interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type emailState struct {
	mailer email.Mailer
}

func (e *emailState) Start(ctx context.Context) error {
	return e.mailer.Start(ctx)
}

func (e *emailState) Stop(ctx context.Context) error {
	return e.mailer.Stop(ctx)
}

var _ Email = (*emailState)(nil)

func ProvideEmail() (Email, error) {
	dialer := gomail.NewDialer(
		viper.GetString("smtp.host"),
		viper.GetInt("smtp.port"),
		viper.GetString("smtp.username"),
		viper.GetString("smtp.password"),
	)

	config := email.Config{
		DefaultFrom: viper.GetString("email.from"),
	}

	sender, err := email.NewGomailSender(dialer, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create gomail sender: %w", err)
	}

	return &emailState{mailer: sender}, nil
}
