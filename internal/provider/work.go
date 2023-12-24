package provider

import (
	"context"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type Work interface {
	Pool() *work.WorkerPool
	Start(context.Context) error
	Stop(context.Context) error
}

type WorkContext struct{}

func ProvideWork() (Work, error) {
	redisPool := redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379", redis.DialPassword(""))
		},
	}

	pool := work.NewWorkerPool(WorkContext{}, 10, "my_app_namespace", &redisPool)

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
