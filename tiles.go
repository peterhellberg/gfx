package gfx

import "image"

// Tiles is a slice of images.
type Tiles []*image.NRGBA

// NewTiles returns a new Tiles from an image.Image and given stride.
func NewTiles(m image.Image, stride int) Tiles {
	var t Tiles

	w := m.Bounds().Dx()
	h := m.Bounds().Dy()

	for y := 0; y < h; y += stride {
		for x := 0; x < w; x += stride {
			dst := NewImage(stride, stride)

			DrawOver(dst, dst.Bounds(), m, Pt(x, y))

			t = append(t, dst)
		}
	}

	return t
}

// NewTiledImage returns a new tiled image.
func NewTiledImage(s Tiles, cols int, layers ...TileLayer) *image.NRGBA {
	if len(s) == 0 || len(layers) == 0 || cols > len(layers[0]) {
		panic("bad NewTiledImage input")
	}

	size := s[0].Bounds().Dx()

	w := cols * size
	h := len(layers[0]) / cols * size

	dst := NewImage(w, h)

	for _, layer := range layers {
		rows := len(layer) / cols

		w = cols * size
		h = rows * size

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				i, x, y := (row*cols)+col, col*size, row*size

				DrawOver(dst, IR(x, y, x+16, y+16), s[layer[i]], ZP)
			}
		}
	}

	return dst
}

// TileLayer is a layer of tiles.
type TileLayer []uint8
