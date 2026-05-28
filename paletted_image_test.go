package gfx

import (
	"image"
	"image/color"
	"testing"
)

func TestNewPaletted(t *testing.T) {
	m := NewPaletted(32, 32, PaletteEN4, PaletteEN4[1])

	if got, want := len(m.Pix), 1024; got != want {
		t.Fatalf("len(m.Pix) = %d, want %d", got, want)
	}
}

func TestNewPalettedImage(t *testing.T) {
	m := NewPalettedImage(IR(0, 0, 32, 32), PaletteEN4)

	if got, want := len(m.Pix), 1024; got != want {
		t.Fatalf("len(m.Pix) = %d, want %d", got, want)
	}
}

func TestNewResizedPalettedImage(t *testing.T) {
	src := NewPaletted(16, 16, PaletteEN4)

	m := NewResizedPalettedImage(src, 32, 32)

	if got, want := len(m.Pix), 1024; got != want {
		t.Fatalf("len(m.Pix) = %d, want %d", got, want)
	}
}

func TestNewScaledPalettedImage(t *testing.T) {
	src := NewPaletted(16, 16, PaletteEN4)

	m := NewScaledPalettedImage(src, 2)

	if got, want := len(m.Pix), 1024; got != want {
		t.Fatalf("len(m.Pix) = %d, want %d", got, want)
	}
}

func TestPalettedColorModel(t *testing.T) {
	m := NewPaletted(16, 16, PaletteEN4)

	c := ColorNRGBA(255, 0, 0, 255)

	got := m.ColorModel().Convert(c).(color.NRGBA)
	want := ColorNRGBA(229, 176, 131, 255)

	if got != want {
		t.Fatalf("m.ColorModel().Convert(c) = %v, want %v", got, want)
	}
}

func TestPalettedPixels(t *testing.T) {
	m := NewPaletted(16, 16, PaletteEN4, PaletteEN4[2])

	if got, want := len(m.Pixels()), 1024; got != want {
		t.Fatalf("len(m.Pixels()) = %d, want %d", got, want)
	}
}

func TestPalettedSubImage(t *testing.T) {
	m := NewPaletted(16, 16, PaletteEN4)

	m.Put(1, 1, 2)

	sm := m.SubImage(IR(1, 1, 4, 4)).(*Paletted)

	if got, want := sm.Bounds(), IR(1, 1, 4, 4); !got.Eq(want) {
		t.Fatalf("sm.Bounds() = %v, want %v", got, want)
	}

	// The sub-image shares storage with the parent. Reading at (1, 1)
	// — the same coordinate Put was called with — must return 2.
	if got, want := sm.Index(1, 1), uint8(2); got != want {
		t.Fatalf("sm.Index(1,1) = %d, want %d", got, want)
	}
}

func TestPalettedOutOfBoundsAccess(t *testing.T) {
	m := NewPalettedImage(IR(10, 20, 14, 24), PaletteEN4)

	m.SetColorIndex(11, 21, 2)

	// Reads outside the bounds return 0 instead of panicking.
	for _, p := range []image.Point{
		{0, 0},   // negative offset (left and above)
		{9, 19},  // one step before Min
		{14, 24}, // exactly at Max (exclusive)
		{50, 50}, // far past Max
	} {
		if got := m.ColorIndexAt(p.X, p.Y); got != 0 {
			t.Fatalf("ColorIndexAt(%v) = %d outside bounds, want 0", p, got)
		}
	}

	// Writes outside the bounds are no-ops instead of panicking.
	for _, p := range []image.Point{
		{0, 0},
		{9, 19},
		{14, 24},
		{50, 50},
	} {
		m.SetColorIndex(p.X, p.Y, 3)
		m.Put(p.X, p.Y, 3)
	}

	// The in-bounds write from before should be untouched.
	if got := m.ColorIndexAt(11, 21); got != 2 {
		t.Fatalf("ColorIndexAt(11, 21) = %d, want 2 (in-bounds pixel disturbed)", got)
	}
}

func TestPalettedNonOriginRect(t *testing.T) {
	m := NewPalettedImage(IR(10, 20, 14, 24), PaletteEN4)

	m.SetColorIndex(10, 20, 1) // top-left
	m.SetColorIndex(13, 23, 3) // bottom-right
	m.SetColorIndex(11, 22, 2) // somewhere in the middle

	for _, tc := range []struct {
		x, y int
		want uint8
	}{
		{10, 20, 1},
		{13, 23, 3},
		{11, 22, 2},
		{12, 21, 0},
	} {
		if got := m.ColorIndexAt(tc.x, tc.y); got != tc.want {
			t.Fatalf("ColorIndexAt(%d, %d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestPalettedOpaque(t *testing.T) {
	t.Run("Opaque", func(t *testing.T) {
		m := NewPaletted(16, 16, PaletteEN4, PaletteEN4[1])

		if !m.Opaque() {
			t.Fatalf("expected image to be opaque")
		}
	})

	t.Run("Not Opaque", func(t *testing.T) {
		m := NewPaletted(16, 16, append(PaletteEN4, ColorTransparent), ColorTransparent)

		if m.Opaque() {
			t.Fatalf("expected image to not be opaque")
		}
	})
}
