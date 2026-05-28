// Package gfxhttp provides HTTP helpers for fetching images, tilesets,
// and map tiles using the gfx package.
//
// It lives in a subpackage so that importers of github.com/peterhellberg/gfx
// who do not need HTTP support are not forced to pull in net/http and its
// transitive dependencies (crypto/tls, crypto/x509, net, etc.).
package gfxhttp

import (
	"context"
	"fmt"
	"image"
	"maps"
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

// StatusError is returned by GetPNG, GetImage and GetTileset when the
// HTTP response has a non-2xx status code. It carries the requested URL
// along with the status line and status code so callers can branch on
// the failure mode.
type StatusError struct {
	URL    string
	Status string
	Code   int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("gfxhttp: GET %s: %s", e.URL, e.Status)
}

// Client performs HTTP requests for the gfx package.
//
// A Client is safe for concurrent use after construction. Option funcs
// must not be applied to a Client that is already in use by another
// goroutine.
type Client struct {
	httpClient *http.Client
	userAgent  string
	headers    http.Header
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

// WithHTTPClient replaces the underlying *http.Client. Passing a nil
// *http.Client is a no-op so that callers can compose options without
// having to nil-check beforehand.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithHeader adds a header that is sent on every request the Client
// makes. Repeated calls with the same key replace the previous value.
// User-Agent set via WithHeader is overridden by WithUserAgent (and by
// the default user agent if neither option is supplied).
func WithHeader(key, value string) Option {
	return func(c *Client) {
		if c.headers == nil {
			c.headers = http.Header{}
		}
		c.headers.Set(key, value)
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

// Get performs an HTTP GET request. The response status is not checked;
// callers are responsible for inspecting resp.StatusCode and closing
// resp.Body. The convenience methods GetPNG, GetImage and GetTileset
// do check the status and return a *StatusError on non-2xx responses.
func (c *Client) Get(ctx context.Context, rawurl string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawurl, nil)
	if err != nil {
		return nil, err
	}

	maps.Copy(req.Header, c.headers)
	req.Header.Set("User-Agent", c.userAgent)

	return c.httpClient.Do(req)
}

// GetPNG retrieves and decodes a remote PNG.
func (c *Client) GetPNG(ctx context.Context, rawurl string) (image.Image, error) {
	resp, err := c.get(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return gfx.DecodePNG(resp.Body)
}

// GetImage retrieves and decodes a remote image. The image format
// (GIF / JPEG / PNG) is auto-detected via image.Decode; the caller is
// responsible for importing the relevant decoders.
func (c *Client) GetImage(ctx context.Context, rawurl string) (image.Image, error) {
	resp, err := c.get(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return gfx.DecodeImage(resp.Body)
}

// GetTileset retrieves a remote PNG and turns it into a *gfx.Tileset.
func (c *Client) GetTileset(ctx context.Context, p gfx.Palette, tileSize image.Point, rawurl string) (*gfx.Tileset, error) {
	m, err := c.GetPNG(ctx, rawurl)
	if err != nil {
		return nil, err
	}

	return gfx.NewTilesetFromImage(p, tileSize, m), nil
}

// get performs Get and verifies that the response status is 2xx,
// closing the body and returning a *StatusError otherwise.
func (c *Client) get(ctx context.Context, rawurl string) (*http.Response, error) {
	resp, err := c.Get(ctx, rawurl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		resp.Body.Close()
		return nil, &StatusError{URL: rawurl, Status: resp.Status, Code: resp.StatusCode}
	}

	return resp, nil
}

// Get performs an HTTP GET request using the Default client.
func Get(ctx context.Context, rawurl string) (*http.Response, error) {
	return Default.Get(ctx, rawurl)
}

// GetPNG retrieves and decodes a remote PNG using the Default client.
func GetPNG(ctx context.Context, rawurl string) (image.Image, error) {
	return Default.GetPNG(ctx, rawurl)
}

// GetImage retrieves and decodes a remote image using the Default client.
func GetImage(ctx context.Context, rawurl string) (image.Image, error) {
	return Default.GetImage(ctx, rawurl)
}

// GetTileset retrieves a remote tileset using the Default client.
func GetTileset(ctx context.Context, p gfx.Palette, tileSize image.Point, rawurl string) (*gfx.Tileset, error) {
	return Default.GetTileset(ctx, p, tileSize, rawurl)
}
