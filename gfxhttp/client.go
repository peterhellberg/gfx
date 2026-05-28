// Package gfxhttp provides HTTP helpers for fetching images, tilesets,
// and map tiles using the gfx package.
//
// It lives in a subpackage so that importers of github.com/peterhellberg/gfx
// who do not need HTTP support are not forced to pull in net/http and its
// transitive dependencies (crypto/tls, crypto/x509, net, etc.).
package gfxhttp

import (
	"image"
	"net/http"
	"time"

	"github.com/peterhellberg/gfx"
)

// DefaultTimeout is the timeout used by clients created with NewClient
// when no explicit timeout has been configured.
const DefaultTimeout = 30 * time.Second

// DefaultUserAgent is the User-Agent header sent by clients created with
// NewClient when no explicit user agent has been configured.
const DefaultUserAgent = "gfx.gfxhttp.Client"

// Client performs HTTP requests for the gfx package.
type Client struct {
	httpClient *http.Client
	userAgent  string
}

// Option configures a Client.
type Option func(*Client)

// WithTimeout sets the request timeout on the underlying *http.Client.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// WithUserAgent sets the User-Agent header sent on each request.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// WithHTTPClient replaces the underlying *http.Client. Use this to share
// a client across packages, install custom transports, etc.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// NewClient creates a new Client with the given options applied.
// A fresh *http.Client with DefaultTimeout is used unless WithHTTPClient
// is provided.
func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: DefaultTimeout},
		userAgent:  DefaultUserAgent,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Default is the package-level Client used by the Get, GetPNG, GetImage
// and GetTileset top-level functions.
var Default = NewClient()

// Get performs an HTTP GET request.
func (c *Client) Get(rawurl string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	return c.httpClient.Do(req)
}

// GetPNG retrieves and decodes a remote PNG.
func (c *Client) GetPNG(rawurl string) (image.Image, error) {
	resp, err := c.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return gfx.DecodePNG(resp.Body)
}

// GetImage retrieves and decodes a remote image. The image format
// (GIF / JPEG / PNG) is auto-detected via image.Decode; the caller is
// responsible for importing the relevant decoders.
func (c *Client) GetImage(rawurl string) (image.Image, error) {
	resp, err := c.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return gfx.DecodeImage(resp.Body)
}

// GetTileset retrieves a remote PNG and turns it into a *gfx.Tileset.
func (c *Client) GetTileset(p gfx.Palette, tileSize image.Point, rawurl string) (*gfx.Tileset, error) {
	m, err := c.GetPNG(rawurl)
	if err != nil {
		return nil, err
	}

	return gfx.NewTilesetFromImage(p, tileSize, m), nil
}

// Get performs an HTTP GET request using the Default client.
func Get(rawurl string) (*http.Response, error) { return Default.Get(rawurl) }

// GetPNG retrieves and decodes a remote PNG using the Default client.
func GetPNG(rawurl string) (image.Image, error) { return Default.GetPNG(rawurl) }

// GetImage retrieves and decodes a remote image using the Default client.
func GetImage(rawurl string) (image.Image, error) { return Default.GetImage(rawurl) }

// GetTileset retrieves a remote tileset using the Default client.
func GetTileset(p gfx.Palette, tileSize image.Point, rawurl string) (*gfx.Tileset, error) {
	return Default.GetTileset(p, tileSize, rawurl)
}
