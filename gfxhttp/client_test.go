package gfxhttp

import (
	"context"
	"errors"
	"image"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/peterhellberg/gfx"
)

func TestGetPNG(t *testing.T) {
	ts := testServer(palettedPNGHandler)
	defer ts.Close()

	m, err := GetPNG(t.Context(), ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	p := m.(image.PalettedImage)

	if got, want := p.Bounds(), gfx.IR(0, 0, 6, 6); !got.Eq(want) {
		t.Fatalf("p.Bounds() = %v, want %v", got, want)
	}

	for _, tc := range []struct {
		x int
		y int
		i uint8
	}{
		{0, 0, 0},
		{2, 4, 6},
	} {
		if got, want := p.ColorIndexAt(tc.x, tc.y), tc.i; got != want {
			t.Fatalf("p.ColorIndexAt(%d, %d) = %v, want %v", tc.x, tc.y, got, want)
		}
	}
}

func TestGetTileset(t *testing.T) {
	ts := testServer(palettedPNGHandler)
	defer ts.Close()

	tileset, err := GetTileset(t.Context(), gfx.PaletteAmmo8, gfx.Pt(3, 3), ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := len(tileset.Tiles), 4; got != want {
		t.Fatalf("len(tileset.Tiles) = %d, want %d", got, want)
	}
}

func TestGetPNGStatusError(t *testing.T) {
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusNotFound)
	})
	defer ts.Close()

	_, err := GetPNG(t.Context(), ts.URL)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	var se *StatusError
	if !errors.As(err, &se) {
		t.Fatalf("expected *StatusError, got %T: %v", err, err)
	}
	if se.Code != http.StatusNotFound {
		t.Fatalf("StatusError.Code = %d, want %d", se.Code, http.StatusNotFound)
	}
	if se.URL != ts.URL {
		t.Fatalf("StatusError.URL = %q, want %q", se.URL, ts.URL)
	}
}

func TestGetContextCancellation(t *testing.T) {
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	})
	defer ts.Close()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	_, err := GetPNG(ctx, ts.URL)
	if err == nil {
		t.Fatalf("expected error from cancelled context, got nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}

func TestWithHeader(t *testing.T) {
	var seenAuth, seenAccept string
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		seenAuth = r.Header.Get("Authorization")
		seenAccept = r.Header.Get("Accept")
		w.Write(palettedPNGData)
	})
	defer ts.Close()

	c := NewClient(
		WithHeader("Authorization", "Bearer secret"),
		WithHeader("Accept", "image/png"),
	)

	if _, err := c.GetPNG(t.Context(), ts.URL); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seenAuth != "Bearer secret" {
		t.Fatalf("Authorization = %q, want %q", seenAuth, "Bearer secret")
	}
	if seenAccept != "image/png" {
		t.Fatalf("Accept = %q, want %q", seenAccept, "image/png")
	}
}

func TestWithUserAgentBeatsWithHeader(t *testing.T) {
	var seenUA string
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		seenUA = r.Header.Get("User-Agent")
		w.Write(palettedPNGData)
	})
	defer ts.Close()

	c := NewClient(
		WithHeader("User-Agent", "via-header"),
		WithUserAgent("via-option"),
	)

	if _, err := c.GetPNG(t.Context(), ts.URL); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seenUA != "via-option" {
		t.Fatalf("User-Agent = %q, want %q", seenUA, "via-option")
	}
}

func TestWithHeaderReplaces(t *testing.T) {
	var seen string
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		seen = r.Header.Get("X-Token")
		w.Write(palettedPNGData)
	})
	defer ts.Close()

	c := NewClient(
		WithHeader("X-Token", "first"),
		WithHeader("X-Token", "second"),
	)

	if _, err := c.GetPNG(t.Context(), ts.URL); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seen != "second" {
		t.Fatalf("X-Token = %q, want %q", seen, "second")
	}
}

func TestClientGetRawResponse(t *testing.T) {
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "value")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	defer ts.Close()

	resp, err := Get(t.Context(), ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if got := resp.Header.Get("X-Test"); got != "value" {
		t.Fatalf("X-Test header = %q, want %q", got, "value")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if string(body) != "hello" {
		t.Fatalf("body = %q, want %q", body, "hello")
	}
}

func TestClientGetRawDoesNotErrOnNon2xx(t *testing.T) {
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	})
	defer ts.Close()

	resp, err := Get(t.Context(), ts.URL)
	if err != nil {
		t.Fatalf("Get returned an error on 500: %v (raw Get should leave status checks to the caller)", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
}

func TestGetImage(t *testing.T) {
	ts := testServer(palettedPNGHandler)
	defer ts.Close()

	m, err := GetImage(t.Context(), ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := m.Bounds(), gfx.IR(0, 0, 6, 6); !got.Eq(want) {
		t.Fatalf("m.Bounds() = %v, want %v", got, want)
	}
}

func TestWithHTTPClientNilIsNoOp(t *testing.T) {
	c := NewClient(WithHTTPClient(nil))

	if c.httpClient == nil {
		t.Fatalf("httpClient is nil after WithHTTPClient(nil); expected the default *http.Client to be kept")
	}
	if got, want := c.httpClient.Timeout, DefaultTimeout; got != want {
		t.Fatalf("httpClient.Timeout = %v, want %v (default kept)", got, want)
	}
}

func TestTileServerDrawTile(t *testing.T) {
	var requested string
	ts := testServer(func(w http.ResponseWriter, r *http.Request) {
		requested = r.URL.Path
		w.Write(palettedPNGData)
	})
	defer ts.Close()

	// Pretend the URL pattern is "<server>/<zoom>/<x>/<y>"; Rawurl uses
	// the gt's Zoom, X, Y in that order.
	server := &TileServer{Format: ts.URL + "/%d/%d/%d", Client: NewClient()}

	dst := gfx.NewImage(16, 16)
	gt := gfx.GT(3, 4, 5)
	gp := gfx.GP(0, 0)

	if err := server.DrawTile(t.Context(), dst, gt, gp); err != nil {
		t.Fatalf("DrawTile: %v", err)
	}

	if want := "/3/4/5"; requested != want {
		t.Fatalf("requested path = %q, want %q", requested, want)
	}
}

func TestNewClientOptions(t *testing.T) {
	c := NewClient(WithUserAgent("ua"), WithTimeout(0))

	if got, want := c.userAgent, "ua"; got != want {
		t.Fatalf("userAgent = %q, want %q", got, want)
	}
	if got, want := c.httpClient.Timeout, time.Duration(0); got != want {
		t.Fatalf("Timeout = %v, want %v", got, want)
	}
}

func testServer(hf http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(hf)
}

func palettedPNGHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(palettedPNGData)
}

var palettedPNGData = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
	0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x06, 0x04, 0x03, 0x00, 0x00, 0x00, 0x12, 0xe2, 0xf2,
	0x7b, 0x00, 0x00, 0x00, 0x24, 0x50, 0x4c, 0x54, 0x45, 0x18, 0x14, 0x25, 0x00, 0x99, 0xdb, 0x12,
	0x4e, 0x89, 0x3e, 0x89, 0x48, 0xe4, 0x3b, 0x44, 0x26, 0x5c, 0x42, 0xfe, 0xae, 0x34, 0xf7, 0x76,
	0x22, 0x2c, 0xe8, 0xf5, 0x63, 0xc7, 0x4d, 0xff, 0xff, 0xff, 0x19, 0x3c, 0x3e, 0xc5, 0xc7, 0x7e,
	0x7a, 0x00, 0x00, 0x00, 0x20, 0x49, 0x44, 0x41, 0x54, 0x08, 0xd7, 0x63, 0x60, 0xe0, 0xb4, 0x64,
	0xe0, 0x62, 0x0e, 0x66, 0x60, 0x60, 0xdd, 0xca, 0xe0, 0x9e, 0x21, 0xc4, 0xe0, 0x9e, 0xa8, 0xc8,
	0xe0, 0x9e, 0x24, 0x01, 0x00, 0x23, 0xc5, 0x03, 0xa8, 0x32, 0xa9, 0x7a, 0x6f, 0x00, 0x00, 0x00,
	0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}
