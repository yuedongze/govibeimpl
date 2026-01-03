package gdrive

import "context"

//go:generate govibeimpl -name GoogleDriveDownloader

// GoogleDriveDownloader downloads files from Google Drive
type GoogleDriveDownloader interface {
	// Initialize instantiates the download with the right credentials.
	Initialize(context.Context, ...any) error
	// Download takes in specification of the file to be downloaded, downloads the file, and returns its content as []byte.
	Download(...any) ([]byte, error)
}
