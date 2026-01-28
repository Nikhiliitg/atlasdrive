package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type FolderCache struct {
	client *redis.Client
}

func NewFolderCache(client *redis.Client) *FolderCache {
	return &FolderCache{client: client}
}

func (c *FolderCache) Get(
	ctx context.Context,
	key string,
	dest interface{},
) (bool, error) {

	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, json.Unmarshal([]byte(val), dest)
}

func (c *FolderCache) Set(
	ctx context.Context,
	key string,
	value interface{},
) error {

	bytes, _ := json.Marshal(value)
	return c.client.Set(ctx, key, bytes, time.Minute).Err()
}

func (c *FolderCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
