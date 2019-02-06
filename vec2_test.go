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

func ExampleVecLerp() {
	a, b := V(1, 2), V(30, 40)

	Dump(
		a.Lerp(b, 0),
		a.Lerp(b, 0.1),
		a.Lerp(b, 0.5),
		a.Lerp(b, 0.9),
		a.Lerp(b, 1),
	)

	// Output:
	//Vec(1, 2)
	//Vec(3.9, 5.8)
	//Vec(15.5, 21)
	//Vec(27.1, 36.2)
	//Vec(30, 40)
}
