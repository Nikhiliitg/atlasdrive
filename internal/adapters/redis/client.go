package redis


import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func Ping(client *redis.Client) error {
	return client.Ping(context.Background()).Err()
}