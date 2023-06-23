package example

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gocraft/work"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Config struct {
	GormDB      *gorm.DB
	MongoClient *mongo.Client
}

type Example struct {
	gormDB      *gorm.DB
	mongoClient *mongo.Client
}

func (e *Example) FiberRoute(router fiber.Router) {
	//
}

func (e *Example) WatermillRoute(pub message.Publisher, sub message.Subscriber, router *message.Router) {
	//
}

func (e *Example) WorkRoute(wp *work.WorkerPool) {
	//
}

func handleCreateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func handleGetUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func handleTriggerEventHello(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func listenEventHello(db *gorm.DB) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		return nil
	}
}

func handleTriggerJobHello(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func doJobHello(db *gorm.DB) work.GenericHandler {
	return func(j *work.Job) error {
		return nil
	}
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(Example), "*"),
	wire.FieldsOf(new(Config), "GormDB", "MongoClient"),
)
