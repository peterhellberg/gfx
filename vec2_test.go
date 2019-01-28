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

func ExampleLerp() {
	a, b := V(1, 2), V(30, 40)

	Dump(
		Lerp(a, b, 0),
		Lerp(a, b, 0.1),
		Lerp(a, b, 0.5),
		Lerp(a, b, 0.9),
		Lerp(a, b, 1),
	)

	// Output:
	//Vec(1, 2)
	//Vec(3.9, 5.8)
	//Vec(15.5, 21)
	//Vec(27.1, 36.2)
	//Vec(30, 40)
}

func ExampleClamp() {
	Dump(
		Clamp(-5, 10, 10),
		Clamp(15, 10, 15),
		Clamp(25, 10, 20),
	)

	// Output:
	// 10
	// 15
	// 20
}
