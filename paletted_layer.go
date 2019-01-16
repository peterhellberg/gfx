package gfx

import (
	"image"
	"image/color"
)

// PalettedLayer represents a layer of paletted tiles.
type PalettedLayer struct {
	*PalettedTileset
	Width int // Width of the layer in number of tiles.
	Data  LayerData
}

// NewPalettedLayer creates a new layer.
func NewPalettedLayer(tileset *PalettedTileset, width int, data LayerData) *PalettedLayer {
	return &PalettedLayer{tileset, width, data}
}

// At returns the color at (x, y).
func (l *PalettedLayer) At(x, y int) color.Color {
	if x < 0 || y < 0 {
		return color.Transparent
	}

	t := l.TileAt(x, y)

	if t != nil {
		return t.At(x%l.Size.X, y%l.Size.Y)
	}

	return color.Transparent
}

// Bounds returns the bounds of the paletted layer.
func (l *PalettedLayer) Bounds() image.Rectangle {
	lpix := len(l.Data)

	switch {
	case l.Width < 1, lpix == 0,
		l.PalettedTileset == nil,
		l.Size.X < 1, l.Size.Y < 1:
		return ZR
	case lpix < l.Width:
		return IR(0, 0, l.Width, 1)
	}

	s := l.Data.Size(l.Width)

	w := s.X * l.Size.X
	h := s.Y * l.Size.Y

	return IR(0, 0, w, h)
}

// ColorModel returns the color model for the paletted layer.
func (l *PalettedLayer) ColorModel() color.Model {
	return color.NRGBAModel
}

// ColorIndexAt returns the palette index of the pixel at (x, y).
func (l *PalettedLayer) ColorIndexAt(x, y int) uint8 {
	t := l.TileAt(x, y)

	if t == nil {
		return 0
	}

	return t.ColorIndexAt(x, y)
}

// TileAt returns the tile image at (x, y).
func (l *PalettedLayer) TileAt(x, y int) image.PalettedImage {
	i := l.indexAt(x, y)

	if i < 0 || i >= len(l.Images) {
		return nil
	}

	return l.Images[i]
}

// TileSize returns the tileset tile size.
func (l *PalettedLayer) TileSize() image.Point {
	return l.Size
}

func (l *PalettedLayer) indexAt(x, y int) int {
	gx, gy := l.gridXY(x, y)

	i := gy*l.Width + gx

	if i < 0 || i >= len(l.Data) {
		return -1
	}

	return l.Data[i]
}

func (l *PalettedLayer) gridXY(x, y int) (int, int) {
	var gx int

	if x >= l.Size.X {
		gx = x / l.Size.X
	}

	var gy int

	if y >= l.Size.Y {
		gy = y / l.Size.Y
	}

	return gx, gy
}
