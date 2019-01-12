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
