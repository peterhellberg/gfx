package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestPt(t *testing.T) {
	got, want := Pt(1, 2), image.Point{1, 2}

	if got != want {
		t.Fatalf("Pt(1,2) = %v, want %v", got, want)
	}
}

func TestIR(t *testing.T) {
	x0, y0, x1, y1 := 10, 10, 30, 30

	got := IR(x0, y0, x1, y1)
	want := IR(x0, y0, x1, y1)

	if got != want {
		t.Fatalf("IR(%d, %d, %d, %d) = %v, want %v", x0, y0, x1, y1, got, want)
	}
}

// TestMixMatchesDrawOver verifies that Mix produces the same pixel as
// draw.Draw with draw.Over for a range of source/destination alpha
// combinations, so the inlined blend stays bit-equivalent to the stdlib.
func TestMixMatchesDrawOver(t *testing.T) {
	cases := []color.RGBA{
		{255, 0, 0, 255},
		{0, 255, 0, 128},
		{0, 0, 255, 64},
		{200, 100, 50, 200},
		{12, 34, 56, 1},
		{0, 0, 0, 0},
	}

	for _, dstColor := range cases {
		for _, srcColor := range cases {
			got := image.NewRGBA(image.Rect(0, 0, 1, 1))
			got.SetRGBA(0, 0, dstColor)
			Mix(got, 0, 0, srcColor)

			want := image.NewRGBA(image.Rect(0, 0, 1, 1))
			want.SetRGBA(0, 0, dstColor)
			draw.Draw(want, want.Bounds(), image.NewUniform(srcColor), image.Point{}, draw.Over)

			if got.RGBAAt(0, 0) != want.RGBAAt(0, 0) {
				t.Fatalf("Mix(dst=%v, src=%v) = %v, want %v",
					dstColor, srcColor, got.RGBAAt(0, 0), want.RGBAAt(0, 0))
			}
		}
	}
}
