package gfx

import "testing"

func TestV(t *testing.T) {
	for _, tc := range []struct {
		x    float64
		y    float64
		want Vec
	}{
		{123, 456, Vec{123, 456}},
		{1.1, 2.2, Vec{1.1, 2.2}},
	} {
		v := V(tc.x, tc.y)

		if v.X != tc.want.X || v.Y != tc.want.Y {
			t.Fatalf("unexpected vector: %v", v)
		}
	}
}
