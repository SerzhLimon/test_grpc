package cashe

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const (
	saveTime = 60 * time.Second
)

type RedisCashe struct {
	client *redis.Client
}

func NewRedisCache() *RedisCashe {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisCashe{
		client: client,
	}
}

func (c RedisCashe) Get(videoID string) ([]byte, error) {
	result, err := c.client.Get(context.Background(), videoID).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New("empty response")
		}
		return nil, err
	}
	return result, nil
}

func (c RedisCashe) Set(videoID string, image []byte) error {
	return c.client.Set(context.Background(), videoID, image, saveTime).Err()
}

func (c RedisCashe) GetClient() *redis.Client {
	return c.client
}

// пример рантайм кэша

/*
type CasheMap struct {
	Storage map[string][]byte
}

func NewStorage() *CasheMap {
	storage := make(map[string][]byte)
	return &CasheMap{
		Storage: storage,
	}
}

func (c *CasheMap) Get(videoID string) ([]byte, error) {
	value, exist := c.Storage[videoID]
	if !exist {
		err := errors.New("empty response")
		return value, err
	}
	return value, nil
}

func (c *CasheMap) Set(videoID string, image []byte) error {
	c.Storage[videoID] = image
	return nil
}
*/
