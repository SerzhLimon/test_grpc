package tests

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	uc "github.com/SerzhLimon/test_grpc/app/internal/usecase"
)

type MockCashe struct {
	GetFunc func(key string) ([]byte, error)
	SetFunc func(key string, value []byte) error
}

func (m *MockCashe) Get(key string) ([]byte, error) {
	if m.GetFunc != nil {
		return m.GetFunc(key)
	}
	return nil, errors.New("not implemented")
}

func (m *MockCashe) Set(key string, value []byte) error {
	if m.SetFunc != nil {
		return m.SetFunc(key, value)
	}
	return errors.New("not implemented")
}

func TestGetPreviewImage(t *testing.T) {
	t.Run("should download image and save to cashe", func(t *testing.T) {
		mockCashe := &MockCashe{
			GetFunc: func(key string) ([]byte, error) {
				return nil, errors.New("not found")
			},
			SetFunc: func(key string, value []byte) error {
				assert.Equal(t, "lKrVuufVMXA", key)
				return nil
			},
		}

		usecase := uc.NewUsecase(mockCashe)
		image, err := usecase.GetPreviewImage("https://www.youtube.com/watch?v=lKrVuufVMXA")

		file, err := os.Open("test_image1/preview.jpg")
		defer file.Close()
		imageTest, err := io.ReadAll(file)

		assert.NoError(t, err)
		assert.Equal(t, imageTest, image)
	})

	t.Run("should download image from cache", func(t *testing.T) {
		file, err := os.Open("test_image2/preview.jpg")
		defer file.Close()
		imageTest, err := io.ReadAll(file)

		mockCashe := &MockCashe{
			GetFunc: func(key string) ([]byte, error) {
				return imageTest, nil
			},
			SetFunc: func(key string, value []byte) error {
				return nil
			},
		}

		usecase := uc.NewUsecase(mockCashe)
		image, err := usecase.GetPreviewImage("https://www.youtube.com/watch?v=lKrVuufVMXA")

		assert.NoError(t, err)
		assert.Equal(t, imageTest, image)
	})
}
