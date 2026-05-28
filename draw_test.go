//go:build !tinygo

package gfx

import "testing"

func TestDrawOver(t *testing.T) {
	src := NewImage(3, 3)

	src.SetRGBA(1, 1, ColorRGBA(75, 0, 130, 64))

	dst := NewImage(3, 3, ColorMagenta)

	DrawOver(dst, dst.Bounds(), src, ZP)
}

func TestDrawSrc(t *testing.T) {
	src := NewImage(3, 3, ColorRGBA(0, 255, 0, 64))
	dst := NewImage(3, 3, ColorMagenta)

	DrawSrc(dst, dst.Bounds(), src, ZP)
}

func TestDrawOverPalettedImage(t *testing.T) {
	src := newTestLayer()
	dst := NewPaletted(8, 16, PaletteEN4)

	DrawPalettedImage(dst, dst.Bounds(), src)
}

func TestDrawLayerOverPaletted(t *testing.T) {
	src := newTestLayer()
	dst := NewPaletted(8, 16, PaletteEN4)

	DrawPalettedLayer(dst, dst.Bounds(), src)
}

func TestDrawPalettedImageNonOriginRect(t *testing.T) {
	src := NewPaletted(8, 16, PaletteEN4)
	for i := range src.Pix {
		src.Pix[i] = 1
	}

	dst := NewPaletted(8, 16, PaletteEN4)
	r := IR(2, 4, 6, 10)

	DrawPalettedImage(dst, r, src)

	for y := 0; y < 16; y++ {
		for x := 0; x < 8; x++ {
			inside := x >= r.Min.X && x < r.Max.X && y >= r.Min.Y && y < r.Max.Y
			got := dst.ColorIndexAt(x, y)
			want := uint8(0)
			if inside {
				want = 1
			}
			if got != want {
				t.Fatalf("dst.ColorIndexAt(%d, %d) = %d, want %d (inside=%v)", x, y, got, want, inside)
			}
		}
	}
}

func TestDrawPalettedLayerNonOriginRect(t *testing.T) {
	src := newTestLayer()
	dst := NewPaletted(16, 12, PaletteEN4)
	r := IR(4, 4, 12, 8)

	DrawPalettedLayer(dst, r, src)

	for y := 0; y < 12; y++ {
		for x := 0; x < 16; x++ {
			inside := x >= r.Min.X && x < r.Max.X && y >= r.Min.Y && y < r.Max.Y
			got := dst.ColorIndexAt(x, y)
			if inside {
				if want := src.ColorIndexAt(x, y); got != want {
					t.Fatalf("dst.ColorIndexAt(%d, %d) = %d, want %d (from src)", x, y, got, want)
				}
			} else if got != 0 {
				t.Fatalf("dst.ColorIndexAt(%d, %d) = %d, want 0 (outside r)", x, y, got)
			}
		}
	}
}

func TestDrawLine(t *testing.T) {
	dst := NewImage(32, 32)

	DrawLine(dst, V(4, 4), V(24, 12), 2, ColorBlue)
	DrawLine(dst, V(4, 4), V(24, 12), 1, ColorGreen)
}

func TestDrawColor(t *testing.T) {
	dst := NewImage(32, 32)

	DrawColor(dst, IR(5, 5, 15, 20), ColorGreen)
}

func TestDrawPolygon(t *testing.T) {
	dst := NewImage(32, 32, ColorBlack)

	p := Polygon{
		{0, 0},
		{20, 2},
		{25, 20},
		{10, 15},
	}

	DrawPolygon(dst, p, 0, ColorMagenta)
	DrawPolygon(dst, p, 1, ColorYellow)
}

func TestDrawPolyline(t *testing.T) {
	dst := NewImage(32, 32, ColorBlack)

	DrawPolyline(dst, Polyline{
		{{0, 0}, {10, 0}, {10, 10}},
		{{10, 10}, {20, 8}, {25, 20}},
	}, 0, ColorMagenta)
}

func TestDrawCircle(t *testing.T) {
	dst := NewImage(32, 32)

	DrawCircle(dst, V(16, 16), 12, 0, ColorMagenta)
	DrawCircle(dst, V(16, 16), 8, 2, ColorYellow)
}

func TestDrawFilledCircle(t *testing.T) {
	dst := NewImage(32, 32)

	DrawCircleFilled(dst, V(16, 16), 8, ColorMagenta)
}

func TestDrawCircleFast(t *testing.T) {
	dst := NewImage(32, 32)

	DrawCircleFast(dst, V(16, 16), 8, ColorMagenta)
}

func TestDrawPointCircle(t *testing.T) {
	dst := NewImage(32, 32)

	DrawPointCircle(dst, Pt(16, 16), 16, 0, ColorMagenta)
	DrawPointCircle(dst, Pt(16, 16), 8, 4, ColorYellow)
}

func ExampleDrawCircle_filled() {
	dst := NewPaletted(15, 13, Palette1Bit, ColorWhite)

	DrawCircle(dst, V(7, 6), 6, 0, ColorBlack)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			if dst.Index(x, y) == 0 {
				Printf("▓▓")
			} else {
				Printf("░░")
			}
		}
		Printf("\n")
	}

	// Output:
	//
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	//
}

func ExampleDrawCircle_annular() {
	dst := NewPaletted(15, 13, Palette1Bit, ColorWhite)

	DrawCircle(dst, V(7, 6), 6, 3, ColorBlack)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			if dst.Index(x, y) == 0 {
				Printf("▓▓")
			} else {
				Printf("░░")
			}
		}
		Printf("\n")
	}

	// Output:
	//
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░
	// ░░░░▓▓▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓░░░░░░░░░░▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░
	// ░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░
	// ░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	//
}
