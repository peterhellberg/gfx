package gfx

import (
	"image"
	"image/color"
	"image/draw"
)

// PalettedImage is an in-memory image of uint8 indices into a given palette.
type PalettedImage struct {
	// Pix holds the image's pixels, as palette indices. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
	// Palette is the image's palette.
	Palette Palette
}

// NewPaletted returns a new paletted image with the given width, height and palette.
func NewPaletted(w, h int, p Palette) *PalettedImage {
	return NewPalettedImage(image.Rect(0, 0, w, h), p)
}

// NewPalettedImage returns a new paletted image with the given bounds and palette.
func NewPalettedImage(r image.Rectangle, p Palette) *PalettedImage {
	w, h := r.Dx(), r.Dy()

	pix := make([]uint8, 1*w*h)

	return &PalettedImage{pix, 1 * w, r, p}
}

// ColorModel returns the color model of the paletted image.
func (p *PalettedImage) ColorModel() color.Model {
	return p.Palette
}

// Bounds returns the bounds of the paletted image.
func (p *PalettedImage) Bounds() image.Rectangle {
	return p.Rect
}

// At retrieves the color at (x, y).
func (p *PalettedImage) At(x, y int) color.Color {
	if len(p.Palette) == 0 {
		return nil
	}

	if !(image.Point{x, y}.In(p.Rect)) {
		return p.Palette[0]
	}

	i := p.PixOffset(x, y)

	return p.Palette[p.Pix[i]]
}

// PixOffset returns the index of the first element of Pix
// that corresponds to the pixel at (x, y).
func (p *PalettedImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*1
}

// Set changes the color at (x, y).
func (p *PalettedImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}

	i := p.PixOffset(x, y)

	p.Pix[i] = uint8(p.Palette.Index(c))
}

// Index returns the color index at (x, y). (Short for ColorIndexAt)
func (p *PalettedImage) Index(x, y int) uint8 {
	return p.ColorIndexAt(x, y)
}

// Put changes the color index at (x, y). (Short for SetColorIndex)
func (p *PalettedImage) Put(x, y int, index uint8) {
	p.SetColorIndex(x, y, index)
}

// ColorIndexAt returns the color index at (x, y).
func (p *PalettedImage) ColorIndexAt(x, y int) uint8 {
	if !(image.Point{x, y}.In(p.Rect)) {
		return 0
	}

	i := p.PixOffset(x, y)

	return p.Pix[i]
}

// SetColorIndex changes the color index at (x, y).
func (p *PalettedImage) SetColorIndex(x, y int, index uint8) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}

	i := p.PixOffset(x, y)

	p.Pix[i] = index
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *PalettedImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)

	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &PalettedImage{
			Palette: p.Palette,
		}
	}

	i := p.PixOffset(r.Min.X, r.Min.Y)

	return &PalettedImage{
		Pix:     p.Pix[i:],
		Stride:  p.Stride,
		Rect:    p.Rect.Intersect(r),
		Palette: p.Palette,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *PalettedImage) Opaque() bool {
	var present [256]bool

	i0, i1 := 0, p.Rect.Dx()

	for y := p.Rect.Min.Y; y < p.Rect.Max.Y; y++ {
		for _, c := range p.Pix[i0:i1] {
			present[c] = true
		}

		i0 += p.Stride
		i1 += p.Stride
	}

	for i, c := range p.Palette {
		if !present[i] {
			continue
		}

		_, _, _, a := c.RGBA()

		if a != 0xffff {
			return false
		}
	}

	return true
}

// Make sure that *PalettedImage implements these interfaces
var (
	_ image.PalettedImage = &PalettedImage{}
	_ draw.Image          = &PalettedImage{}
)
