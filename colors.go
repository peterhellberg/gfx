package gfx

import "image/color"

// Standard colors
var (
	ColorBlack       = color.Black
	ColorWhite       = color.White
	ColorTransparent = color.Transparent
	ColorOpaque      = color.Opaque
	ColorRed         = Color3Bit(1)
	ColorGreen       = Color3Bit(2)
	ColorBlue        = Color3Bit(3)
	ColorCyan        = Color3Bit(4)
	ColorMagenta     = Color3Bit(5)
	ColorYellow      = Color3Bit(6)
)

// ColorWithAlpha creates a new color.NRGBA based
// on the provided color.Color and alpha arguments.
func ColorWithAlpha(c color.Color, a uint8) color.NRGBA {
	nc := color.NRGBAModel.Convert(c).(color.NRGBA)

	nc.A = a

	return nc
}

// ColorNRGBA constructs a color.NRGBA
func ColorNRGBA(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{r, g, b, a}
}
