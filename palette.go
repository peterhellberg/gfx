package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

// Palette is a slice of colors.
type Palette []color.NRGBA

// Color returns the color at intex n.
func (p Palette) Color(n int) color.NRGBA {
	if n >= 0 && n < p.Len() {
		return p[n]
	}

	return color.NRGBA{}
}

// Len returns the number of colors in the palette
func (p Palette) Len() int {
	return len(p)
}

// Random color from the palette.
func (p Palette) Random() color.NRGBA {
	return p[rand.Intn(p.Len())]
}

// Image returns a new image based on the input image, but with colors from the palette.
func (p Palette) Image(src image.Image) *PalettedImage {
	dst := NewPalettedImage(src.Bounds(), p)

	draw.Draw(dst, dst.Bounds(), src, image.ZP, draw.Src)

	return dst
}

// Convert returns the palette color closest to c in Euclidean R,G,B space.
func (p Palette) Convert(c color.Color) color.Color {
	if len(p) == 0 {
		return color.NRGBA{}
	}

	return p[p.Index(c)]
}

// Index returns the index of the palette color closest to c in Euclidean
// R,G,B,A space.
func (p Palette) Index(c color.Color) int {
	cr, cg, cb, ca := c.RGBA()
	ret, bestSum := 0, uint32(1<<32-1)

	for i, v := range p {
		vr, vg, vb, va := v.RGBA()
		sum := sqDiff(cr, vr) + sqDiff(cg, vg) + sqDiff(cb, vb) + sqDiff(ca, va)

		if sum < bestSum {
			if sum == 0 {
				return i
			}

			ret, bestSum = i, sum
		}
	}

	return ret
}

// sqDiff returns the squared-difference of x and y, shifted by 2 so that
// adding four of those won't overflow a uint32.
//
// x and y are both assumed to be in the range [0, 0xffff].
func sqDiff(x, y uint32) uint32 {
	d := x - y

	return (d * d) >> 2
}
