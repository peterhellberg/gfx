package gfx

import (
	"image/color"
	"testing"
)

func TestColorNRGBA(t *testing.T) {
	got := ColorNRGBA(11, 22, 33, 44)
	want := color.NRGBA{11, 22, 33, 44}

	if got != want {
		t.Fatalf("NRGBA(11,22,33,44) = %v, want %v", got, want)
	}
}

func TestColorPICO8(t *testing.T) {
	for n, want := range PalettePICO8 {
		if got := ColorPICO8(n); got != want {
			t.Fatalf("ColorPICO8(%d) = %v, want %v", n, got, want)
		}
	}
}

func TestColorEDG64(t *testing.T) {
	for n, want := range PaletteEDG64 {
		if got := ColorEDG64(n); got != want {
			t.Fatalf("ColorEDG64(%d) = %v, want %v", n, got, want)
		}
	}
}

func TestColorEDG36(t *testing.T) {
	for n, want := range PaletteEDG36 {
		if got := ColorEDG36(n); got != want {
			t.Fatalf("ColorEDG36(%d) = %v, want %v", n, got, want)
		}
	}
}

func TestColorEDG32(t *testing.T) {
	for n, want := range PaletteEDG32 {
		if got := ColorEDG32(n); got != want {
			t.Fatalf("ColorEDG32(%d) = %v, want %v", n, got, want)
		}
	}
}

func TestColorEDG16(t *testing.T) {
	for n, want := range PaletteEDG16 {
		if got := ColorEDG16(n); got != want {
			t.Fatalf("ColorEDG16(%d) = %v, want %v", n, got, want)
		}
	}
}

func TestColorCGA(t *testing.T) {
	for n, want := range PaletteCGA {
		if got := ColorCGA(n); got != want {
			t.Fatalf("ColorCGA(%d) = %v, want %v", n, got, want)
		}
	}
}

func TestColor2BitGrayScale(t *testing.T) {
	for n, want := range Palette2BitGrayScale {
		if got := Color2BitGrayScale(n); got != want {
			t.Fatalf("Color2BitGrayScale(%d) = %v, want %v", n, got, want)
		}
	}
}
