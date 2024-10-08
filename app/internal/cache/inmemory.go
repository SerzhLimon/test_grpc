package cache

import "github.com/pkg/errors"

type cacheMap struct {
	Storage map[string][]byte
}

func NewStorage() *cacheMap {
	storage := make(map[string][]byte)
	return &cacheMap{
		Storage: storage,
	}
}

func (c cacheMap) Get(videoID string) ([]byte, error) {
	value, exist := c.Storage[videoID]
	if !exist {
		err := errors.New("not found")
		return value, err
	}
	return value, nil
}

func (c cacheMap) Set(videoID string, image []byte) error {
	c.Storage[videoID] = image
	return nil
}
