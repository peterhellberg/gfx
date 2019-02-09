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

// BlockColor contains a Light, Medium and Dark color.
type BlockColor struct {
	Light  color.NRGBA
	Medium color.NRGBA
	Dark   color.NRGBA
}

// Block colors based on PaletteTango.
var (
	BlockColorYellow = BlockColor{PaletteTango[0], PaletteTango[1], PaletteTango[2]}
	BlockColorOrange = BlockColor{PaletteTango[3], PaletteTango[4], PaletteTango[5]}
	BlockColorBrown  = BlockColor{PaletteTango[6], PaletteTango[7], PaletteTango[8]}
	BlockColorGreen  = BlockColor{PaletteTango[9], PaletteTango[10], PaletteTango[11]}
	BlockColorBlue   = BlockColor{PaletteTango[12], PaletteTango[13], PaletteTango[14]}
	BlockColorPurple = BlockColor{PaletteTango[15], PaletteTango[16], PaletteTango[17]}
	BlockColorRed    = BlockColor{PaletteTango[18], PaletteTango[19], PaletteTango[20]}
	BlockColorWhite  = BlockColor{PaletteTango[21], PaletteTango[22], PaletteTango[23]}
	BlockColorBlack  = BlockColor{PaletteTango[24], PaletteTango[25], PaletteTango[26]}

	// BlockColors is a slice of all the default block colors.
	BlockColors = []BlockColor{
		BlockColorYellow,
		BlockColorOrange,
		BlockColorBrown,
		BlockColorGreen,
		BlockColorBlue,
		BlockColorPurple,
		BlockColorRed,
		BlockColorWhite,
		BlockColorBlack,
	}
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

// ColorRGBA constructs a color.RGBA.
func ColorRGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

// LerpColors performs linear interpolation between two colors.
func LerpColors(c0, c1 color.Color, t float64) color.Color {
	switch {
	case t <= 0:
		return c0
	case t >= 1:
		return c1
	}

	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c1.RGBA()

	fr0, fg0, fb0, fa0 := float64(r0), float64(g0), float64(b0), float64(a0)
	fr1, fg1, fb1, fa1 := float64(r1), float64(g1), float64(b1), float64(a1)

	return color.RGBA64{
		uint16(Clamp(fr0+(fr1-fr0)*t, 0, 0xffff)),
		uint16(Clamp(fg0+(fg1-fg0)*t, 0, 0xffff)),
		uint16(Clamp(fb0+(fb1-fb0)*t, 0, 0xffff)),
		uint16(Clamp(fa0+(fa1-fa0)*t, 0, 0xffff)),
	}
}
