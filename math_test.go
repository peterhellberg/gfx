package gfx

import "testing"

func TestClampInt(t *testing.T) {
	cases := []struct {
		x, lo, hi, want int
	}{
		{-5, 0, 10, 0},
		{5, 0, 10, 5},
		{15, 0, 10, 10},
	}
	for _, tc := range cases {
		if got := Clamp(tc.x, tc.lo, tc.hi); got != tc.want {
			t.Fatalf("Clamp(%d, %d, %d) = %d, want %d", tc.x, tc.lo, tc.hi, got, tc.want)
		}
	}
}

func TestSignInt(t *testing.T) {
	if got := Sign(-3); got != -1 {
		t.Fatalf("Sign(-3) = %d, want -1", got)
	}
	if got := Sign(0); got != 0 {
		t.Fatalf("Sign(0) = %d, want 0", got)
	}
	if got := Sign(7); got != 1 {
		t.Fatalf("Sign(7) = %d, want 1", got)
	}
}

func TestLerpInt(t *testing.T) {
	if got := Lerp(0, 10, 0.5); got != 5 {
		t.Fatalf("Lerp(0, 10, 0.5) = %d, want 5", got)
	}
	if got := Lerp(0, 100, 0.25); got != 25 {
		t.Fatalf("Lerp(0, 100, 0.25) = %d, want 25", got)
	}
}

func ExampleMathMin() {
	Dump(
		MathMin(-1, 1),
		MathMin(1, 2),
		MathMin(3, 2),
	)

	// Output:
	// -1
	// 1
	// 2
}

func ExampleMathMax() {
	Dump(
		MathMax(-1, 1),
		MathMax(1, 2),
		MathMax(3, 2),
	)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleMathAbs() {
	Dump(
		MathAbs(-2),
		MathAbs(-1),
		MathAbs(0),
		MathAbs(1),
		MathAbs(2),
	)

	// Output:
	// 2
	// 1
	// 0
	// 1
	// 2
}

func ExampleMathSqrt() {
	Dump(
		MathSqrt(1),
		MathSqrt(2),
		MathSqrt(3),
	)

	// Output:
	// 1
	// 1.4142135623730951
	// 1.7320508075688772
}

func ExampleMathSin() {
	Dump(
		MathSin(1),
		MathSin(2),
		MathSin(3),
	)

	// Output:
	// 0.8414709848078965
	// 0.9092974268256816
	// 0.1411200080598672
}

func ExampleMathCos() {
	Dump(
		MathCos(1),
		MathCos(2),
		MathCos(3),
	)

	// Output:
	// 0.5403023058681398
	// -0.4161468365471424
	// -0.9899924966004454
}

func ExampleMathCeil() {
	Dump(
		MathCeil(0.2),
		MathCeil(1.4),
		MathCeil(2.6),
	)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleMathFloor() {
	Dump(
		MathFloor(0.2),
		MathFloor(1.4),
		MathFloor(2.6),
	)

	// Output:
	// 0
	// 1
	// 2
}

func ExampleMathHypot() {
	Dump(
		MathHypot(15, 8),
		MathHypot(5, 12),
		MathHypot(3, 4),
	)

	// Output:
	// 17
	// 13
	// 5
}

func ExampleSign() {
	Dump(
		Sign(-2),
		Sign(0),
		Sign(2),
	)

	// Output:
	// -1
	// 0
	// 1
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

func ExampleLerp() {
	Dump(
		Lerp(0.0, 2.0, 0.1),
		Lerp(1.0, 10.0, 0.5),
		Lerp(2.0, 4.0, 0.5),
	)

	// Output:
	// 0.2
	// 5.5
	// 3
}
