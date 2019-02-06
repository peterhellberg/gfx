package gfx

import (
	"image"
	"math"
	"testing"
)

func TestV(t *testing.T) {
	for _, tc := range []struct {
		x    float64
		y    float64
		want Vec
	}{
		{123, 456, Vec{123, 456}},
		{1.1, 2.2, Vec{1.1, 2.2}},
	} {
		if got := V(tc.x, tc.y); got != tc.want {
			t.Fatalf("unexpected vector: %v", got)
		}
	}
}

func TestIV(t *testing.T) {
	for _, tc := range []struct {
		x    int
		y    int
		want Vec
	}{
		{123, 456, Vec{123, 456}},
		{789, 333, Vec{789, 333}},
	} {
		if got := IV(tc.x, tc.y); got != tc.want {
			t.Fatalf("unexpected vector: %v", got)
		}
	}
}

func TestPV(t *testing.T) {
	for _, tc := range []struct {
		p    image.Point
		want Vec
	}{
		{Pt(123, 456), Vec{123, 456}},
		{Pt(789, 333), Vec{789, 333}},
	} {
		if got := PV(tc.p); got != tc.want {
			t.Fatalf("unexpected vector: %v", got)
		}
	}
}

func TestUnit(t *testing.T) {
	for angle, want := range map[float64]Vec{
		1: V(0.5403023058681398, 0.8414709848078965),
		2: V(-0.4161468365471424, 0.9092974268256816),
		9: V(-0.9111302618846769, 0.4121184852417566),
	} {
		if got := Unit(angle); got != want {
			t.Fatalf("Unit(%v) = %v, want %v", angle, got, want)
		}
	}
}

func TestVecEq(t *testing.T) {
	for _, tc := range []struct {
		u    Vec
		v    Vec
		want bool
	}{
		{V(1, 1), V(1, 1), true},
		{V(1, 1), V(2, 2), false},
	} {
		if got := tc.u.Eq(tc.v); got != tc.want {
			t.Fatalf("%v.Eq(%v) = %v, want %v", tc.u, tc.v, got, tc.want)
		}
	}
}

func ExampleVecAdd() {
	Dump(
		V(1, 1).Add(V(2, 3)),
		V(3, 3).Add(V(-1, -2)),
	)

	// Output:
	// gfx.V(3, 4)
	// gfx.V(2, 1)
}

func ExampleVecAddXY() {
	Dump(
		V(1, 1).AddXY(2, 3),
		V(3, 3).AddXY(-1, -2),
	)

	// Output:
	// gfx.V(3, 4)
	// gfx.V(2, 1)
}

func ExampleVecSub() {
	Dump(
		V(1, 1).Sub(V(2, 3)),
		V(3, 3).Sub(V(-1, -2)),
	)

	// Output:
	// gfx.V(-1, -2)
	// gfx.V(4, 5)
}

func ExampleVecTo() {
	Dump(
		V(1, 1).To(V(2, 3)),
		V(3, 3).To(V(-1, -2)),
	)

	// Output:
	// gfx.V(1, 2)
	// gfx.V(-4, -5)
}

func ExampleVecMod() {
	Dump(
		V(1, 1).Mod(V(2.5, 3)),
		V(2, 5.5).Mod(V(2, 3)),
	)

	// Output:
	// gfx.V(1, 1)
	// gfx.V(0, 2.5)
}

func ExampleVecAbs() {
	Dump(
		V(1, -1).Abs(),
		V(-2, -2).Abs(),
		V(3, 6).Abs(),
	)

	// Output:
	// gfx.V(1, 1)
	// gfx.V(2, 2)
	// gfx.V(3, 6)
}

func ExampleVecMax() {
	Dump(
		V(1, 1).Max(V(2.5, 3)),
		V(2, 5.5).Max(V(2, 3)),
	)

	// Output:
	// gfx.V(2.5, 3)
	// gfx.V(2, 5.5)
}

func ExampleVecMin() {
	Dump(
		V(1, 1).Min(V(2.5, 3)),
		V(2, 5.5).Min(V(2, 3)),
	)

	// Output:
	// gfx.V(1, 1)
	// gfx.V(2, 3)
}

func ExampleVecDot() {
	Dump(
		V(1, 1).Dot(V(2.5, 3)),
		V(2, 5.5).Dot(V(2, 3)),
	)

	// Output:
	// 5.5
	// 20.5
}

func ExampleVecCross() {
	Dump(
		V(1, 1).Cross(V(2.5, 3)),
		V(2, 5.5).Cross(V(2, 3)),
	)

	// Output:
	// 0.5
	// -5
}

func ExampleVecProject() {
	Dump(
		V(1, 1).Project(V(2.5, 3)),
		V(2, 5.5).Project(V(2, 3)),
	)

	// Output:
	// gfx.V(0.9016393442622948, 1.0819672131147537)
	// gfx.V(3.153846153846153, 4.73076923076923)
}

func ExampleVecMap() {
	Dump(
		V(1.1, 1).Map(math.Ceil),
		V(1.1, 2.5).Map(math.Round),
	)

	// Output:
	// gfx.V(2, 1)
	// gfx.V(1, 3)
}

func ExampleRect() {
	Dump(
		V(10, 10).Rect(-1, -2, 3, 4),
		V(3, 4).Rect(1.5, 2.2, 3.3, 4.5),
	)

	// Output:
	// gfx.R(9, 8, 13, 14)
	// gfx.R(4.5, 6.2, 6.3, 8.5)
}

func ExampleBounds() {
	Dump(
		V(10, 10).Bounds(-1, -2, 3, 4),
		V(3, 4).Bounds(1.5, 2.2, 3.3, 4.5),
	)

	// Output:
	// (9,8)-(13,14)
	// (4,6)-(6,8)
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
	// gfx.V(1, 2)
	// gfx.V(3.9, 5.8)
	// gfx.V(15.5, 21)
	// gfx.V(27.1, 36.2)
	// gfx.V(30, 40)
}

func ExampleCentroid() {

	Dump(
		Centroid(V(1, 1), V(6, 1), V(3, 4)),
		Centroid(V(0, 0), V(10, 0), V(5, 10)),
	)

	// Output:
	// gfx.V(3.3333333333333335, 2)
	// gfx.V(5, 3.3333333333333335)
}
