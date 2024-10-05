package usecase

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Usecase struct{}

func (u *Usecase) GetPreviewImage(url string) ([]byte, error) {
	var result []byte
	videoID, err := u.extractVideoID(url)
	if err != nil {
		return result, fmt.Errorf("invalid YouTube URL")
	}
	result, err = u.downloadImage(fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", videoID))
	if err != nil {
		return result, fmt.Errorf("failed to download preview image")
	}
	return result, nil
}

func (u *Usecase) GetPreviewImageSlice(urls []string) ([][]byte, error) {
	var result [][]byte
	for _, v := range urls {
		res, err := u.GetPreviewImage(v)
		if err != nil {
			return result, fmt.Errorf("failed to download preview image")
		}
		result = append(result, res)
	}

	return result, nil
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
