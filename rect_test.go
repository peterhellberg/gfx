package gfx

import "testing"

func TestNewRect(t *testing.T) {
	for _, tc := range []struct {
		min Vec
		max Vec
	}{
		{V(1, 2), V(3, 4)},
		{V(5, 6), V(7, 8)},
	} {
		r := NewRect(tc.min, tc.max)

		if r.Min != tc.min || r.Max != tc.max {
			t.Fatalf("unexpected rect: %v", r)
		}
	}
}

func TestRectOverlaps(t *testing.T) {
	for _, tc := range []struct {
		r1   Rect
		r2   Rect
		want bool
	}{
		{R(10, 10, 25, 25), R(20, 20, 30, 30), true},
		{R(10, 10, 20, 20), R(30, 30, 40, 40), false},
	} {
		if got := tc.r1.Overlaps(tc.r2); got != tc.want {
			t.Fatalf("%v.Overlaps(%v) = %v, want %v", tc.r1, tc.r2, got, tc.want)
		}
	}
}
