package provider

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type Watermill interface {
	Router() *message.Router
	Publisher() message.Publisher
	Subscriber() message.Subscriber
	Serve(ctx context.Context) error
	Clean(ctx context.Context) error
}

type watermillState struct {
	router     *message.Router
	publisher  message.Publisher
	subscriber message.Subscriber
}

func ProvideWatermill(ctx context.Context) (Watermill, error) {
	var (
		sub message.Subscriber
		pub message.Publisher
	)

	logger := watermill.NewStdLogger(false, false)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	pubSubChan := gochannel.NewGoChannel(gochannel.Config{}, logger)

	sub = pubSubChan
	pub = pubSubChan

	return &watermillState{
		router:     router,
		publisher:  pub,
		subscriber: sub,
	}, nil
}

func (w *watermillState) Router() *message.Router {
	return w.router
}

func (w *watermillState) Publisher() message.Publisher {
	return w.publisher
}

func (w *watermillState) Subscriber() message.Subscriber {
	return w.subscriber
}

func (w *watermillState) Serve(ctx context.Context) error {
	return w.router.Run(ctx)
}

func (w *watermillState) Clean(ctx context.Context) error {
	if w.router.IsRunning() {
		if err := w.router.Close(); err != nil {
			return err
		}
	}

	if err := w.subscriber.Close(); err != nil {
		return err
	}

	if err := w.publisher.Close(); err != nil {
		return err
	}

	return nil
}
