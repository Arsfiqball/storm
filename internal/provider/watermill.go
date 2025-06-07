package provider

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
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

func ProvideWatermill(ctx context.Context, sl Slog) (Watermill, error) {
	var (
		sub message.Subscriber
		pub message.Publisher
	)

	logger := watermill.NewSlogLogger(sl.Logger())

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	// Check if Redis is configured
	redisURL := viper.GetString("redis.url")
	if redisURL == "" {
		// Fallback to gochannel if Redis is not configured
		pubSubChan := gochannel.NewGoChannel(gochannel.Config{}, logger)
		sub = pubSubChan
		pub = pubSubChan
	} else {
		// Parse Redis URL
		redisOptions, err := redis.ParseURL(redisURL)
		if err != nil {
			return nil, err
		}

		// Get additional Redis configuration from viper
		consumerGroup := viper.GetString("redis.consumer_group")
		if consumerGroup == "" {
			consumerGroup = "storm-consumer-group"
		}

		consumerPrefix := viper.GetString("redis.consumer_prefix")
		if consumerPrefix == "" {
			consumerPrefix = "storm-consumer"
		}

		// Create Redis client
		redisClient := redis.NewClient(redisOptions)

		// Test connection
		if _, err := redisClient.Ping(ctx).Result(); err != nil {
			return nil, err
		}

		// Configure Redis publisher
		pub, err = redisstream.NewPublisher(
			redisstream.PublisherConfig{
				Client: redisClient,
			},
			logger,
		)
		if err != nil {
			return nil, err
		}

		// Configure Redis subscriber with a consumer group
		sub, err = redisstream.NewSubscriber(
			redisstream.SubscriberConfig{
				Client:                 redisClient,
				ConsumerGroup:          consumerGroup,
				Consumer:               consumerPrefix + "-" + watermill.NewShortUUID(),
				MaxIdleTime:            viper.GetDuration("redis.max_idle_time"),
				BlockTime:              viper.GetDuration("redis.block_time"),
				ClaimInterval:          viper.GetDuration("redis.claim_interval"),
				ClaimBatchSize:         viper.GetInt64("redis.claim_batch_size"),
				CheckConsumersInterval: viper.GetDuration("redis.check_consumers_interval"),
				ConsumerTimeout:        viper.GetDuration("redis.consumer_timeout"),
				NackResendSleep:        viper.GetDuration("redis.nack_resend_sleep"),
			},
			logger,
		)
		if err != nil {
			return nil, err
		}
	}

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
