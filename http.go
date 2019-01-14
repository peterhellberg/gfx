package gfx

import (
	"image"
	"net/http"
	"time"
)

// DefaultClient is the default HTTP client used by the gfx package.
var DefaultClient = &http.Client{
	Timeout: 30 * time.Second,
}

// GetPNG retrieves a remote PNG using DefaultClient
func GetPNG(rawurl string) (image.Image, error) {
	resp, err := DefaultClient.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return DecodePNG(resp.Body)
}
