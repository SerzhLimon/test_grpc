package cashe

import (
	"github.com/pkg/errors"
)

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
