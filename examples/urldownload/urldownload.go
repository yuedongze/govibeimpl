package urldownload

//go:generate govibeimpl -name URLDownloader

// URLDownloader downloads the content behind an URL via a HTTP GET request
type URLDownloader interface {
	// Download downloads the URL and returns the content as []byte
	Download(...any) ([]byte, error)
}
