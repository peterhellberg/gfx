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
