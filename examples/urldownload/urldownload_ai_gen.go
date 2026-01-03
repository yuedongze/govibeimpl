// AI generated implementation. PROCEED WITH CAUTION.
package urldownload

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// DownloadOptions defines the parameters for a download request.
// It allows passing headers and other configuration to the HTTP request.
//
// Example usage:
//
//	opts := DownloadOptions{
//		URL: "https://api.example.com/data",
//		Headers: map[string]string{
//			"Authorization": "Bearer token",
//			"Accept":        "application/json",
//		},
//	}
//	downloader := &HTTPDownloader{Timeout: 30 * time.Second}
//	content, err := downloader.Download(opts)
type DownloadOptions struct {
	URL     string
	Headers map[string]string
}

// HTTPDownloader is a concrete implementation of the URLDownloader interface.
// It uses the standard library's net/http package to perform requests.
type HTTPDownloader struct {
	// Client is the underlying HTTP client. If nil, http.DefaultClient is used.
	Client *http.Client
	// Timeout sets a timeout for the request if the Client field is nil.
	Timeout time.Duration
}

// Download downloads the content behind a URL via an HTTP GET request.
// It accepts either a single string (the URL) or a DownloadOptions struct.
//
// Example with string:
//
//	downloader := &HTTPDownloader{}
//	data, err := downloader.Download("https://google.com")
//
// Example with DownloadOptions:
//
//	data, err := downloader.Download(DownloadOptions{URL: "https://example.com"})
func (d *HTTPDownloader) Download(args ...any) ([]byte, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("download requires at least one argument (string or DownloadOptions)")
	}

	var url string
	var headers map[string]string

	switch v := args[0].(type) {
	case string:
		url = v
	case DownloadOptions:
		url = v.URL
		headers = v.Headers
	case *DownloadOptions:
		if v != nil {
			url = v.URL
			headers = v.Headers
		}
	default:
		return nil, fmt.Errorf("unsupported argument type: %T", args[0])
	}

	if url == "" {
		return nil, fmt.Errorf("target URL cannot be empty")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := d.Client
	if client == nil {
		if d.Timeout > 0 {
			client = &http.Client{Timeout: d.Timeout}
		} else {
			client = http.DefaultClient
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("received non-2xx status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

var _ URLDownloader = (*HTTPDownloader)(nil)