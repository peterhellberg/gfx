package gfx

// IntAbs returns the absolute value of x.
func IntAbs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

// IntMin returns the smaller of x or y.
func IntMin(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// IntMax returns the larger of x or y.
func IntMax(x, y int) int {
	if x > y {
		return x
	}

	return y
}
