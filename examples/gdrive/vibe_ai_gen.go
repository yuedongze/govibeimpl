package gdrive

import (
	"context"
	"errors"
	"fmt"
	"io"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type googleDriveDownloader struct {
	service *drive.Service
}

// AuthConfig represents the credentials required to initialize the downloader.
type AuthConfig struct {
	CredentialsJSON []byte
	APIKey          string
}

// DownloadParams represents the parameters required to locate and download a file.
type DownloadParams struct {
	FileID string
}

// NewGoogleDriveDownloader returns a new instance of GoogleDriveDownloader.
func NewGoogleDriveDownloader() GoogleDriveDownloader {
	return &googleDriveDownloader{}
}

// Initialize instantiates the download with the right credentials.
// Expects an AuthConfig as the first variadic argument.
func (d *googleDriveDownloader) Initialize(ctx context.Context, args ...any) error {
	if len(args) == 0 {
		return errors.New("initialization requires at least one argument of type AuthConfig")
	}

	cfg, ok := args[0].(AuthConfig)
	if !ok {
		return errors.New("first argument must be of type AuthConfig")
	}

	var opts []option.ClientOption
	if len(cfg.CredentialsJSON) > 0 {
		opts = append(opts, option.WithCredentialsJSON(cfg.CredentialsJSON))
	} else if cfg.APIKey != "" {
		opts = append(opts, option.WithAPIKey(cfg.APIKey))
	} else {
		return errors.New("AuthConfig must provide either CredentialsJSON or APIKey")
	}

	srv, err := drive.NewService(ctx, opts...)
	if err != nil {
		return fmt.Errorf("failed to create Google Drive service: %w", err)
	}

	d.service = srv
	return nil
}

// Download takes in specification of the file to be downloaded and returns its content.
// Expects a DownloadParams as the first variadic argument.
func (d *googleDriveDownloader) Download(args ...any) ([]byte, error) {
	if d.service == nil {
		return nil, errors.New("downloader not initialized; call Initialize first")
	}

	if len(args) == 0 {
		return nil, errors.New("download requires at least one argument of type DownloadParams")
	}

	params, ok := args[0].(DownloadParams)
	if !ok {
		return nil, errors.New("first argument must be of type DownloadParams")
	}

	if params.FileID == "" {
		return nil, errors.New("FileID cannot be empty")
	}

	resp, err := d.service.Files.Get(params.FileID).Download()
	if err != nil {
		return nil, fmt.Errorf("failed to initiate file download: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	return data, nil
}