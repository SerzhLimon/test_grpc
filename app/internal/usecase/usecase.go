package usecase

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/pkg/errors"
)

type Cashe interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

type Usecase struct {
	cashe Cashe
}

func NewUsecase(cashe Cashe) *Usecase {
	return &Usecase{
		cashe: cashe,
	}
}

func (u *Usecase) GetPreviewImage(url string) ([]byte, error) {
	var image []byte
	videoID, err := u.extractVideoID(url)
	if err != nil {
		return image, fmt.Errorf("invalid YouTube URL")
	}
	
	image, err = u.cashe.Get(videoID)
	if err == nil {
		fmt.Println("-------------------------\nreturn image from storage")
		return image, nil
	}

	image, err = u.downloadImage(fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", videoID))
	if err != nil {
		return image, fmt.Errorf("failed to download preview image")
	}

	if err = u.cashe.Set(videoID, image); err != nil {
		log.Println("can't save image in storage", err)
	}

	return image, nil
}

func (u *Usecase) GetPreviewImageSlice(urls []string) ([][]byte, error) {
	var images [][]byte
	var wg sync.WaitGroup
	var m sync.Mutex
	var errChan = make(chan error, len(urls))

	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			img, err := u.GetPreviewImage(url)
			if err != nil {
				errChan <- err
				return
			}
			m.Lock()
			images = append(images, img)
			m.Unlock()
		}(url)
	}
	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, fmt.Errorf("failed to download preview image")
	}

	return images, nil
}

func (u *Usecase) extractVideoID(videoURL string) (string, error) {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return "", errors.Wrap(err, "extractVideoID")
	}
	return parsedURL.Query().Get("v"), nil
}

func (u *Usecase) downloadImage(imageURL string) ([]byte, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}
