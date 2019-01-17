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
		ts.Tiles = append(ts.Tiles, NewTile(p, s.X, td[i]))
	}

	return ts
}

// NewTile returns a new paletted image with the given pix, stride and palette.
func NewTile(p Palette, cols int, pix []uint8) *Paletted {
	return &Paletted{
		Palette: p,
		Stride:  cols,
		Pix:     pix,
		Rect:    calcRect(cols, pix),
	}
}

func calcRect(cols int, pix []uint8) image.Rectangle {
	s := calcSize(cols, pix)

	return IR(0, 0, s.X, s.Y)
}

func calcSize(cols int, pix []uint8) image.Point {
	l := len(pix)

	if l < cols {
		return Pt(cols, 1)
	}

	rows := l / cols

	if rows*cols == l {
		return Pt(cols, rows)
	}

	if rows%cols > 0 {
		rows++
	}

	return Pt(cols, rows)
}
