package urldownload

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPDownloader is an implementation of the URLDownloader interface.
type HTTPDownloader struct {
	Client *http.Client
}

// NewHTTPDownloader creates a new instance with a default timeout.
func NewHTTPDownloader() *HTTPDownloader {
	return &HTTPDownloader{
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Download performs a GET request to the provided URL.
// It expects the first argument in the variadic slice to be a string representing the URL.
func (d *HTTPDownloader) Download(args ...any) ([]byte, error) {
	if len(args) == 0 {
		return nil, errors.New("url argument is missing")
	}

	url, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("expected first argument to be string, got %T", args[0])
	}

	resp, err := d.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return content, nil
}

// Ensure HTTPDownloader implements URLDownloader at compile time.
var _ URLDownloader = (*HTTPDownloader)(nil)