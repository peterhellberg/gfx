package gfx

import "math"

// MathMin returns the smaller of x or y.
func MathMin(x, y float64) float64 {
	return math.Min(x, y)
}

// MathMax returns the larger of x or y.
func MathMax(x, y float64) float64 {
	return math.Max(x, y)
}

// Sign returns -1 for values < 0, 0 for 0, and 1 for values > 0.
func Sign(x float64) float64 {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

// Clamp returns x clamped to the interval [min, max].
//
// If x is less than min, min is returned. If x is more than max, max is returned. Otherwise, x is
// returned.
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Lerp does linear interpolation between two values.
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}
