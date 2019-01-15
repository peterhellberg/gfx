package gfx

import "image"

// Sprites is a slice of images.
type Sprites []*image.NRGBA

// NewSprites returns Sprites from an image.Image and given stride.
func NewSprites(m image.Image, stride int) Sprites {
	var s Sprites

	w := m.Bounds().Dx()
	h := m.Bounds().Dy()

	for y := 0; y < h; y += stride {
		for x := 0; x < w; x += stride {
			dst := NewImage(stride, stride)

			DrawOver(dst, dst.Bounds(), m, Pt(x, y))

			s = append(s, dst)
		}
	}

	return s
}

// NewSpriteImage returns a new sprite image.
func NewSpriteImage(s Sprites, cols int, grid []uint8) *image.NRGBA {
	if len(s) == 0 || cols > len(grid) {
		return nil
	}

	size := s[0].Bounds().Dx()
	rows := len(grid) / cols

	w := cols * size
	h := rows * size

	dst := NewImage(w, h)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			i, x, y := (row*cols)+col, col*size, row*size

			DrawOver(dst, IR(x, y, x+16, y+16), s[grid[i]], ZP)
		}
	}

	return dst
}
