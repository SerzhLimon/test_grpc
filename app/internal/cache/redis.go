package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const (
	saveTime = 60 * time.Second
)

type Rediscache struct {
	client *redis.Client
}

func NewRedisCache() *Rediscache {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &Rediscache{
		client: client,
	}
}

func (c Rediscache) Get(videoID string) ([]byte, error) {
	result, err := c.client.Get(context.Background(), videoID).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return result, nil
}

func (c Rediscache) Set(videoID string, image []byte) error {
	return c.client.Set(context.Background(), videoID, image, saveTime).Err()
}

func (c Rediscache) GetClient() *redis.Client {
	return c.client
}
