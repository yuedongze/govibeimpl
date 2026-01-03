package urldownload

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DefaultURLDownloader is a standard implementation of the URLDownloader interface
type DefaultURLDownloader struct {
	client *http.Client
}

// NewDefaultURLDownloader creates a new instance with a default http client configuration
func NewDefaultURLDownloader() *DefaultURLDownloader {
	return &DefaultURLDownloader{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Download implements the URLDownloader interface.
// It expects the first argument to be a string representing the URL.
func (d *DefaultURLDownloader) Download(args ...any) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("url argument is required")
	}

	url, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("expected first argument to be string, got %T", args[0])
	}

	resp, err := d.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

// Ensure DefaultURLDownloader implements URLDownloader interface
var _ URLDownloader = (*DefaultURLDownloader)(nil)