package usecase

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Usecase interface {
	ExtractVideoID(videoURL string) (string, error)
	DownloadImage(imageURL string) ([]byte, error)
}

func extractVideoID(videoURL string) (string, error) {
	re := regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]{11})`)
	matches := re.FindStringSubmatch(videoURL)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid YouTube URL")
	}
	return matches[1], nil
}

func downloadImage(imageURL string) ([]byte, error) {
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
