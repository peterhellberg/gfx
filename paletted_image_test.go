package gfx

import (
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
