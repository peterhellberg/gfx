package gfx

import (
	"image/color"
	"image/draw"
)

// DrawBresenhamLine draws a line using Bresenham's line algorithm.
//
// http://en.wikipedia.org/wiki/Bresenham's_line_algorithm
func DrawBresenhamLine(m draw.Image, from, to Vec, c color.Color) {
	x0, y0 := from.XY()
	x1, y1 := to.XY()

	abs := func(i float64) float64 {
		if i < 0 {
			return -i
		}

		return i
	}

	steep := abs(y0-y1) > abs(x0-x1)

	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}

	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := abs(y1 - y0)
	e := dx / 2
	y := y0

	var ystep float64 = -1

	if y0 < y1 {
		ystep = 1
	}

	for x := x0; x <= x1; x++ {
		if steep {
			Mix(m, int(y), int(x), c)
		} else {
			Mix(m, int(x), int(y), c)
		}

		e -= dy

		if e < 0 {
			y += ystep
			e += dx
		}
	}
}
