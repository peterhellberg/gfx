package gfx

import (
	"image/color"
	"testing"
)

func TestPalettes(t *testing.T) {
	for _, tc := range []struct {
		Name       string
		Palette    Palette
		ColorFunc  func(int) color.NRGBA
		ColorCount int
	}{
		{"2BitGrayScale", Palette2BitGrayScale, Color2BitGrayScale, 4},
		{"3Bit", Palette3Bit, Color3Bit, 8},
		{"ARQ4", PaletteARQ4, ColorARQ4, 4},
		{"CGA", PaletteCGA, ColorCGA, 16},
		{"EDG16", PaletteEDG16, ColorEDG16, 16},
		{"EDG32", PaletteEDG32, ColorEDG32, 32},
		{"EDG36", PaletteEDG36, ColorEDG36, 36},
		{"EDG64", PaletteEDG64, ColorEDG64, 64},
		{"EDG8", PaletteEDG8, ColorEDG8, 8},
		{"EN4", PaletteEN4, ColorEN4, 4},
		{"Ink", PaletteInk, ColorInk, 5},
		{"PICO8", PalettePICO8, ColorPICO8, 16},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			if got, want := len(tc.Palette), tc.ColorCount; got != want {
				t.Fatalf("unexpected number of colors: %d, want %d", got, want)
			}

			for n, want := range tc.Palette {
				if got := tc.ColorFunc(n); got != want {
					t.Fatalf("ColorFunc(%d) = %v, want %v", n, got, want)
				}
			}
		})
	}
}
