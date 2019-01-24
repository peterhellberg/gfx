package gfx

import "image/color"

// Standard colors transparent, opaque, black, white, red, green, blue, cyan, magenta, and yellow.
var (
	ColorTransparent = ColorNRGBA(0, 0, 0, 0)
	ColorOpaque      = ColorNRGBA(0xFF, 0xFF, 0xFF, 0xFF)
	ColorBlack       = Palette1Bit.Color(0)
	ColorWhite       = Palette1Bit.Color(1)
	ColorRed         = Palette3Bit.Color(1)
	ColorGreen       = Palette3Bit.Color(2)
	ColorBlue        = Palette3Bit.Color(3)
	ColorCyan        = Palette3Bit.Color(4)
	ColorMagenta     = Palette3Bit.Color(5)
	ColorYellow      = Palette3Bit.Color(6)
)

// ColorWithAlpha creates a new color.NRGBA based
// on the provided color.Color and alpha arguments.
func ColorWithAlpha(c color.Color, a uint8) color.NRGBA {
	nc := color.NRGBAModel.Convert(c).(color.NRGBA)

	nc.A = a

	return nc
}

// ColorNRGBA constructs a color.NRGBA.
func ColorNRGBA(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{r, g, b, a}
}
