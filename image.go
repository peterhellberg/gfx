package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
)

// Pt returns an image.Point for the given x and y.
func Pt(x, y int) image.Point {
	return image.Pt(x, y)
}

// IR returns an image.Rectangle for the given input.
func IR(x0, y0, x1, y1 int) image.Rectangle {
	return image.Rect(x0, y0, x1, y1)
}

// EachPixel calls the provided function for each pixel in the provided rectangle.
func EachPixel(m draw.Image, r image.Rectangle, fn func(x, y int)) {
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

// Mix the current pixel color at the image.Point with the given color.
func Mix(m draw.Image, p image.Point, c color.Color) {
	draw.Draw(m, image.Rect(p.X, p.Y, p.X+1, p.Y+1), image.NewUniform(c), image.ZP, draw.Over)
}

// Set x and y to the given color.
func Set(m draw.Image, x, y int, c color.Color) {
	m.Set(x, y, c)
}

// SetPoint to the given color.
func SetPoint(m draw.Image, p image.Point, c color.Color) {
	m.Set(p.X, p.Y, c)
}

// NewImage creates an image of the given size and color
func NewImage(w, h int, c color.Color) *image.NRGBA {
	m := image.NewNRGBA(image.Rect(0, 0, w, h))

	draw.Draw(m, m.Bounds(), NewUniform(c), image.ZP, draw.Src)

	return m
}

// NewUniform creates a new uniform image of the given color.
func NewUniform(c color.Color) *image.Uniform {
	return image.NewUniform(c)
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
