package provider

import (
	"context"
	"fmt"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type Work interface {
	Pool() *work.WorkerPool
	Start(context.Context) error
	Stop(context.Context) error
}

type WorkContext struct{}

func ProvideWork() (Work, error) {
	// Get Redis URL from Viper config
	redisURL := viper.GetString("redis.url")
	if redisURL == "" {
		return nil, fmt.Errorf("redis URL not configured")
	}

	// Parse Redis URL and create pool
	redisPool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisURL)
		},
	}

	// Test connection to Redis
	conn := redisPool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	pool := work.NewWorkerPool(WorkContext{}, 1, "my_app_namespace", redisPool)

	return &workState{pool: pool}, nil
}

type workState struct {
	pool *work.WorkerPool
}

func (w *workState) Pool() *work.WorkerPool {
	return w.pool
}

func (w *workState) Start(ctx context.Context) error {
	w.pool.Start()

	return nil
}

func (w *workState) Stop(ctx context.Context) error {
	w.pool.Stop()

	return nil
}
