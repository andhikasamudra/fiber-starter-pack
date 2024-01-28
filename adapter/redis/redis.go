package redis

import (
	"github.com/andhikasamudra/fiber-starter-pack/internal/env"
	"github.com/redis/go-redis/v9"
)

type Adapter struct {
	Client *redis.Client
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) GetRedisConnection() {
	conn := redis.NewClient(&redis.Options{
		Addr:     env.RedisURL(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	a.Client = conn
}
