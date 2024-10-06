package cashe

import (
	"github.com/pkg/errors"
)

type CasheMap struct {
	Storage map[string][]byte
}

type Cashe interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

func NewStorage() *CasheMap {
	storage := make(map[string][]byte)
	return &CasheMap{
		Storage: storage,
	}
}

func (c *CasheMap) Get(key string) ([]byte, error) {
	value := c.Storage[key]
	if len(value) == 0 {
		err := errors.New("empty response")
		return value, err
	}
	return value, nil
}

func (c *CasheMap) Set(key string, value []byte) error {
	c.Storage[key] = value
	return nil
}
