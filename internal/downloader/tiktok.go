package downloader

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type TikTokDownloader struct{}

func NewTikTok() *TikTokDownloader {
	return &TikTokDownloader{}
}

func (t TikTokDownloader) Download(fileName string) (io.ReadCloser, error) {

	downloadUrl := fmt.Sprintf("https://tikcdn.io/ssstik/%s", fileName)

	resp, err := http.Get(downloadUrl)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (t TikTokDownloader) ParseMessage(url string) (string, error) {
	var re = regexp.MustCompile(`https:\/\/www\.tiktok\.com\/.+\/video\/(?P<fileName>\d+)`)

	match := re.FindStringSubmatch(url)

	if len(match) != 2 {
		return "", fmt.Errorf("invalid url")
	}

	return match[1], nil
}
