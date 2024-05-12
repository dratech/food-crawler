package uploader

import (
	"fmt"
	"io"
	"net/http"
)

type Dropbox struct {
	token string
}

func NewDropbox(token string) *Dropbox {
	return &Dropbox{
		token,
	}
}

func (d Dropbox) Upload(src io.ReadCloser, location string) error {
	defer src.Close()

	req, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", src)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", d.token))
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Dropbox-API-Arg", fmt.Sprintf(`{"path":"/%s.mp4"}`, location))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.Status != "200 OK" {
		return fmt.Errorf("Failed to upload file: "+ resp.Status)
	}

	return nil
}