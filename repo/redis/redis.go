package redis

import (
	"context"
	"github.com/WHUPRJ/woj-server/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var _ global.Repo = (*Repo)(nil)

type Repo struct {
	client *redis.Client
	log    *zap.Logger
}

func (r *Repo) Setup(g *global.Global) {
	r.log = g.Log

	r.client = redis.NewClient(&redis.Options{
		Addr:     g.Conf.Redis.Address,
		Password: g.Conf.Redis.Password,
		DB:       g.Conf.Redis.Db,
	})

	_, err := r.client.Ping(context.Background()).Result()
	if err != nil {
		r.log.Fatal("Redis ping failed", zap.Error(err))
		return
	}
}

func (r *Repo) Get() interface{} {
	return r.client
}

func (r *Repo) Close() error {
	return r.client.Close()
}
