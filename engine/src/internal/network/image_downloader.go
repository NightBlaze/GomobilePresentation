package network

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

var (
	imageHttpClientTransport = func() *http.Transport {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 10
		t.MaxConnsPerHost = 10
		t.MaxIdleConnsPerHost = 10
		t.IdleConnTimeout = 10 * time.Second
		t.DisableKeepAlives = false
		return t
	}()

	imageHttpClient = &http.Client{
		Timeout:   15 * time.Second,
		Transport: imageHttpClientTransport,
	}
)

type ImageDownloader struct {
}

func NewImageDownloader() *ImageDownloader {
	return &ImageDownloader{}
}

func (i *ImageDownloader) Download(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	response, err := imageHttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() //#nosec CWE-703

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	return io.ReadAll(response.Body)
}
