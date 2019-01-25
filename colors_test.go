package gfx

import (
	"image/color"
	"testing"
)

func TestColorNRGBA(t *testing.T) {
	got := ColorRGBA(11, 22, 33, 44)
	want := color.RGBA{11, 22, 33, 44}

	if got != want {
		t.Fatalf("RGBA(11,22,33,44) = %v, want %v", got, want)
	}
}
