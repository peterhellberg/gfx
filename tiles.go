package gfx

import "image"

// Tiles is a slice of paletted images.
type Tiles []PalettedImage

// Tileset is a paletted tileset.
type Tileset struct {
	Palette Palette     // Palette of the tileset.
	Size    image.Point // Size is the size of each tile.
	Tiles   Tiles       // Images containst all of the images in the tileset.
}

// TilesetData is the raw data in a tileset
type TilesetData [][]uint8

// NewTileset creates a new paletted tileset.
func NewTileset(p Palette, s image.Point, td TilesetData) *Tileset {
	ts := &Tileset{Palette: p, Size: s}

	for i := 0; i < len(td); i++ {
		ts.Tiles = append(ts.Tiles, NewTile(td[i], s.X, p))
	}

	return ts
}

// NewTile returns a new paletted image with the given pix, stride and palette.
func NewTile(pix []uint8, stride int, p Palette) *Paletted {
	return &Paletted{
		Rect:    IR(0, 0, stride, len(pix)/stride+len(pix)%stride),
		Pix:     pix,
		Stride:  stride,
		Palette: p,
	}
}
