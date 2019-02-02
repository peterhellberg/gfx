package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

// Palette is a slice of colors.
type Palette []color.NRGBA

// Color returns the color at index n.
func (p Palette) Color(n int) color.NRGBA {
	if n >= 0 && n < p.Len() {
		return p[n]
	}

	return color.NRGBA{}
}

// Len returns the number of colors in the palette.
func (p Palette) Len() int {
	return len(p)
}

// Random color from the palette.
func (p Palette) Random() color.NRGBA {
	return p[rand.Intn(p.Len())]
}

// Sheet returns a paletted image with all of the colors in the palette.
func (p Palette) Sheet(width int) *Paletted {
	if width < 1 {
		width = 1
	}

	var pix []uint8

	for i := range p {
		pix = append(pix, uint8(i))
	}

	m := NewTile(p, width, pix)

	return m
}

// Tile returns a new image based on the input image, but with colors from the palette.
func (p Palette) Tile(src image.Image) *Paletted {
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

// AsColorPalette converts the Palette to a color.Palette.
func (p Palette) AsColorPalette() color.Palette {
	var cp = make(color.Palette, len(p))

	for i, c := range p {
		cp[i] = c
	}

	return cp
}

// ComplexAt returns the color at the given complex128 value.
func (p Palette) ComplexAt(z complex128) color.Color {
	t := CmplxPhase(z)/Pi + 1

	if t > 1 {
		t = 2 - t
	}

	return p.At(t)
}

// At returns the color at the given float64 value (range 0-1)
func (p Palette) At(t float64) color.Color {
	n := len(p)
	if t <= 0 || math.IsNaN(t) {
		return p[0]
	}
	if t >= 1 {
		return p[n-1]
	}

	i := int(math.Floor(t * float64(n-1)))
	s := 1 / float64(n-1)

	return lerpColors(p[i], p[i+1], (t-float64(i)*s)/s)
}

func lerpColors(c0, c1 color.Color, t float64) color.Color {
	if t <= 0 {
		return c0
	}

	if t >= 1 {
		return c1
	}

	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c1.RGBA()

	fr0, fg0, fb0, fa0 := float64(r0), float64(g0), float64(b0), float64(a0)
	fr1, fg1, fb1, fa1 := float64(r1), float64(g1), float64(b1), float64(a1)

	fr := Clamp(fr0+(fr1-fr0)*t, 0, 0xffff)
	fg := Clamp(fg0+(fg1-fg0)*t, 0, 0xffff)
	fb := Clamp(fb0+(fb1-fb0)*t, 0, 0xffff)
	fa := Clamp(fa0+(fa1-fa0)*t, 0, 0xffff)

	r := uint16(fr)
	g := uint16(fg)
	b := uint16(fb)
	a := uint16(fa)

	return color.RGBA64{r, g, b, a}
}

// sqDiff returns the squared-difference of x and y, shifted by 2 so that
// adding four of those won't overflow a uint32.
//
// x and y are both assumed to be in the range [0, 0xffff].
func sqDiff(x, y uint32) uint32 {
	d := x - y

	return (d * d) >> 2
}
