package infrastructure

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Connection *redis.Client
	IsClosed   bool
}

func NewRedis(hostName string, password string, dbNum int) *Redis {
	redisWrapper := new(Redis)

	redisWrapper.Connection = redis.NewClient(&redis.Options{
		Addr:     hostName,
		Password: password,
		DB:       dbNum,
	})

	redisWrapper.IsClosed = false
	return redisWrapper
}

func StartConsume(redis Redis, channelName string, ctx context.Context, fn func(string)) {
	pubsub := redis.Connection.Subscribe(ctx, channelName)
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		go fn(msg.Payload)
	}
}
