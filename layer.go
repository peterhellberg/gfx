package gfx

import (
	"image"
	"image/color"
)

// Layer represents a layer of paletted tiles.
type Layer struct {
	Tileset *Tileset
	Width   int // Width of the layer in number of tiles.
	Data    LayerData
}

// LayerData is the data for a layer
type LayerData []int

// Size returns the size of the layer data given the number of columns.
func (ld LayerData) Size(cols int) image.Point {
	l := len(ld)

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

// NewLayer creates a new layer.
func NewLayer(tileset *Tileset, width int, data LayerData) *Layer {
	return &Layer{tileset, width, data}
}

// At returns the color at (x, y).
func (l *Layer) At(x, y int) color.Color {
	if x < 0 || y < 0 {
		return color.Transparent
	}

	t := l.TileAt(x, y)

	if t != nil {
		return t.At(x%l.Tileset.Size.X, y%l.Tileset.Size.Y)
	}

	return color.Transparent
}

// Bounds returns the bounds of the paletted layer.
func (l *Layer) Bounds() image.Rectangle {
	lpix := len(l.Data)

	switch {
	case l.Width < 1, lpix == 0,
		l.Tileset == nil,
		l.Tileset.Size.X < 1, l.Tileset.Size.Y < 1:
		return ZR
	case lpix < l.Width:
		return IR(0, 0, l.Width, 1)
	}

	s := l.Data.Size(l.Width)

	w := s.X * l.Tileset.Size.X
	h := s.Y * l.Tileset.Size.Y

	return IR(0, 0, w, h)
}

// ColorModel returns the color model for the paletted layer.
func (l *Layer) ColorModel() color.Model {
	return color.NRGBAModel
}

// ColorIndexAt returns the palette index of the pixel at (x, y).
func (l *Layer) ColorIndexAt(x, y int) uint8 {
	t := l.TileAt(x, y)

	if t == nil {
		return 0
	}

	return t.ColorIndexAt(x, y)
}

// TileAt returns the tile image at (x, y).
func (l *Layer) TileAt(x, y int) image.PalettedImage {
	i := l.indexAt(x, y)

	if i < 0 || i >= len(l.Tileset.Tiles) {
		return nil
	}

	return l.Tileset.Tiles[i]
}

// TileSize returns the tileset tile size.
func (l *Layer) TileSize() image.Point {
	return l.Tileset.Size
}

// GfxPalette retrieves the layer palette.
func (l *Layer) GfxPalette() Palette {
	return l.Tileset.Palette
}

// ColorPalette retrieves the layer palette.
func (l *Layer) ColorPalette() color.Palette {
	return l.Tileset.Palette.AsColorPalette()
}

func (l *Layer) indexAt(x, y int) int {
	gx, gy := l.gridXY(x, y)

	i := gy*l.Width + gx

	if i < 0 || i >= len(l.Data) {
		return -1
	}

	return l.Data[i]
}

func (l *Layer) gridXY(x, y int) (int, int) {
	var gx int

	if x >= l.Tileset.Size.X {
		gx = x / l.Tileset.Size.X
	}

	var gy int

	if y >= l.Tileset.Size.Y {
		gy = y / l.Tileset.Size.Y
	}

	return gx, gy
}
