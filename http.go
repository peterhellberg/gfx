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

// Get performs a HTTP GET request using the DefaultClient.
func Get(rawurl string) (*http.Response, error) {
	return DefaultClient.Get(rawurl)
}

// GetPNG retrieves a remote PNG using DefaultClient
func GetPNG(rawurl string) (image.Image, error) {
	resp, err := Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return DecodePNG(resp.Body)
}

// GetTileset retrieves a remote tileset using GetPNG.
func GetTileset(p Palette, tileSize image.Point, rawurl string) (*Tileset, error) {
	m, err := GetPNG(rawurl)
	if err != nil {
		return nil, err
	}

	return NewTilesetFromImage(p, tileSize, m), nil
}
