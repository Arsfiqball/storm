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
	FakePublisher() FakeWatermillPublisher
	FakeSubscriber() FakeWatermillSubscriber
	Serve(ctx context.Context) error
	Clean(ctx context.Context) error
}

type watermillState struct {
	router     *message.Router
	publisher  message.Publisher
	subscriber message.Subscriber
	fakePub    FakeWatermillPublisher
	fakeSub    FakeWatermillSubscriber
}

func ProvideWatermill(ctx context.Context, sl Slog) (Watermill, error) {
	var (
		sub     message.Subscriber
		pub     message.Publisher
		fakePub FakeWatermillPublisher
		fakeSub FakeWatermillSubscriber
	)

	logger := watermill.NewSlogLogger(sl.Logger())

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	// Check component test mode
	isComponentTest := viper.GetString("mode") == "component_test"

	// Check if Redis is configured
	redisURL := viper.GetString("redis.url")

	if isComponentTest {
		// Use fake publisher and subscriber in component test mode
		sub = newFakeWatermillSubscriber()
		pub = newFakeWatermillPublisher()
	} else if redisURL == "" {
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
			consumerGroup = "app-consumer-group"
		}

		consumerPrefix := viper.GetString("redis.consumer_prefix")
		if consumerPrefix == "" {
			consumerPrefix = "app-consumer"
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
		fakePub:    fakePub,
		fakeSub:    fakeSub,
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

func (w *watermillState) FakePublisher() FakeWatermillPublisher {
	return w.fakePub
}

func (w *watermillState) FakeSubscriber() FakeWatermillSubscriber {
	return w.fakeSub
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

type FakeWatermillSubscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	Close() error
}

type fakeWatermillSubscriber struct {
	message map[string]chan *message.Message
}

func newFakeWatermillSubscriber() FakeWatermillSubscriber {
	return &fakeWatermillSubscriber{
		message: make(map[string]chan *message.Message),
	}
}

func (ms *fakeWatermillSubscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	if ms.message[topic] == nil {
		ms.message[topic] = make(chan *message.Message)
	}

	return ms.message[topic], nil
}

func (ms *fakeWatermillSubscriber) Close() error {
	for _, ch := range ms.message {
		close(ch)
	}

	return nil
}

type FakeWatermillPublisher interface {
	Publish(topic string, messages ...*message.Message) error
	Close() error
	Reset()
	GetMessages(topic string) []message.Message
}

type fakeWatermillPublisher struct {
	messages map[string][]*message.Message
}

func newFakeWatermillPublisher() FakeWatermillPublisher {
	return &fakeWatermillPublisher{
		messages: make(map[string][]*message.Message),
	}
}

func (mp *fakeWatermillPublisher) Publish(topic string, messages ...*message.Message) error {
	mp.messages[topic] = append(mp.messages[topic], messages...)
	return nil
}

func (mp *fakeWatermillPublisher) Close() error {
	return nil
}

func (mp *fakeWatermillPublisher) Reset() {
	mp.messages = make(map[string][]*message.Message)
}

func (mp *fakeWatermillPublisher) GetMessages(topic string) []message.Message {
	var messages []message.Message

	for _, msg := range mp.messages[topic] {
		messages = append(messages, *msg.Copy())
	}

	return messages
}
