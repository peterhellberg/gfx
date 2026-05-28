package gfxhttp

import (
	"image"
	"image/draw"

	"github.com/peterhellberg/gfx"
)

// TileServer represents a map tile server. Format is a Sprintf format
// string that takes zoom, x, y in that order (matching gfx.GeoTile.Rawurl).
type TileServer struct {
	Format string
	Client *Client
}

// NewTileServer creates a TileServer that uses the Default client.
func NewTileServer(format string) *TileServer {
	return &TileServer{Format: format, Client: Default}
}

// GetImage fetches the image for the given GeoTile.
func (ts *TileServer) GetImage(gt gfx.GeoTile) (image.Image, error) {
	return ts.Client.GetImage(gt.Rawurl(ts.Format))
}

// TileImage fetches the image for the given GeoTile using the Default client.
func TileImage(gt gfx.GeoTile, format string) (image.Image, error) {
	return Default.GetImage(gt.Rawurl(format))
}

// DrawTile fetches and draws the given GeoTile on dst.
func (ts *TileServer) DrawTile(dst draw.Image, gt gfx.GeoTile, gp gfx.GeoPoint) error {
	src, err := ts.GetImage(gt)
	if err != nil {
		return err
	}

	gt.Draw(dst, gp, src)

	return nil
}

// DrawNeighbors fetches and draws the 8 neighbors of the given GeoTile.
func (ts *TileServer) DrawNeighbors(dst draw.Image, gt gfx.GeoTile, gp gfx.GeoPoint) error {
	for _, n := range gt.Neighbors() {
		if err := ts.DrawTile(dst, n, gp); err != nil {
			return err
		}
	}

	return nil
}

// DrawTileAndNeighbors draws the given GeoTile and its 8 neighbors on dst.
// If drawing the central tile fails the neighbors are still drawn, matching
// the prior behavior of gfx.GeoTileServer.
func (ts *TileServer) DrawTileAndNeighbors(dst draw.Image, gt gfx.GeoTile, gp gfx.GeoPoint) error {
	if err := ts.DrawTile(dst, gt, gp); err != nil {
		return nil
	}

	return ts.DrawNeighbors(dst, gt, gp)
}
