package gfx

import (
	"image"
	"image/draw"
)

// ResizeImage using nearest neighbor scaling from src to dst.
func ResizeImage(src image.Image, dst draw.Image) {
	w := dst.Bounds().Dx()
	h := dst.Bounds().Dy()

	xRatio := src.Bounds().Dx()<<16/w + 1
	yRatio := src.Bounds().Dy()<<16/h + 1

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			sx := ((x * xRatio) >> 16)
			sy := ((y * yRatio) >> 16)

			dst.Set(x, y, src.At(sx, sy))
		}
	}
}

// NewResizedImage returns an image with the provided dimensions.
func NewResizedImage(src image.Image, w, h int) *image.NRGBA {
	dst := NewImage(w, h)

	ResizeImage(src, dst)

	return dst
}

// NewScaledImage returns an image scaled by the provided scaling factor.
func NewScaledImage(src image.Image, s float64) *image.NRGBA {
	b := src.Bounds()

	if b.Empty() {
		return &image.NRGBA{}
	}

	return NewResizedImage(src, int(float64(b.Dx())*s), int(float64(b.Dy())*s))
}

// NewResizedPalettedImage returns an image with the provided dimensions.
func NewResizedPalettedImage(src *PalettedImage, w, h int) *PalettedImage {
	dst := NewPaletted(w, h, src.Palette)

	ResizeImage(src, dst)

	return dst
}

// NewScaledPalettedImage returns a paletted image scaled by the provided scaling factor.
func NewScaledPalettedImage(src *PalettedImage, s float64) *PalettedImage {
	b := src.Bounds()

	return NewResizedPalettedImage(src, int(float64(b.Dx())*s), int(float64(b.Dy())*s))
}
