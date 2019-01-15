package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
)

// NewImage creates an image of the given size (optionally filled with a color)
func NewImage(w, h int, colors ...color.Color) *image.NRGBA {
	m := NewNRGBA(image.Rect(0, 0, w, h))

	if len(colors) > 0 {
		draw.Draw(m, m.Bounds(), NewUniform(colors[0]), image.ZP, draw.Src)
	}

	return m
}

// NewNRGBA returns a new NRGBA image with the given bounds.
func NewNRGBA(r image.Rectangle) *image.NRGBA {
	return image.NewNRGBA(r)
}

// NewUniform creates a new uniform image of the given color.
func NewUniform(c color.Color) *image.Uniform {
	return image.NewUniform(c)
}

// Pt returns an image.Point for the given x and y.
func Pt(x, y int) image.Point {
	return image.Pt(x, y)
}

// IR returns an image.Rectangle for the given input.
func IR(x0, y0, x1, y1 int) image.Rectangle {
	return image.Rect(x0, y0, x1, y1)
}

// Mix the current pixel color at x and y with the given color.
func Mix(m draw.Image, x, y int, c color.Color) {
	_, _, _, a := c.RGBA()

	switch a {
	case 0xFFFF:
		m.Set(x, y, c)
	default:
		draw.Draw(m, image.Rect(x, y, x+1, y+1), image.NewUniform(c), image.ZP, draw.Over)
	}
}

// MixPoint the current pixel color at the image.Point with the given color.
func MixPoint(m draw.Image, p image.Point, c color.Color) {
	Mix(m, p.X, p.Y, c)
}

// Set x and y to the given color.
func Set(m draw.Image, x, y int, c color.Color) {
	m.Set(x, y, c)
}

// SetPoint to the given color.
func SetPoint(m draw.Image, p image.Point, c color.Color) {
	m.Set(p.X, p.Y, c)
}

// EachPixel calls the provided function for each pixel in the provided rectangle.
func EachPixel(m image.Image, r image.Rectangle, fn func(x, y int)) {
	r = r.Intersect(m.Bounds())

	if r.Empty() {
		return
	}

	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			fn(x, y)
		}
	}
}

// SavePNG saves an image using the provided file name.
func SavePNG(fn string, m image.Image) error {
	w, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer w.Close()

	return EncodePNG(w, m)
}

// EncodePNG encodes an image as PNG to the provided io.Writer
func EncodePNG(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

// DecodePNG decodes a PNG from the provided io.Reader
func DecodePNG(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}
