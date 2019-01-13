package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Polygon is represented by a list of vectors
type Polygon []Vec

// Bounds return the bounds of the polygon rectangle
func (p Polygon) Bounds() image.Rectangle {
	return p.Rect().Bounds()
}

// Rect is the polygon rectangle
func (p Polygon) Rect() Rect {
	r := R(math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64)

	for _, u := range p {
		x, y := u.XY()

		if x > r.Max.X {
			r.Max.X = x
		}

		if y > r.Max.Y {
			r.Max.Y = y
		}

		if x < r.Min.X {
			r.Min.X = x
		}

		if y < r.Min.Y {
			r.Min.Y = y
		}
	}

	return r
}

// EachPixel calls the provided function for each pixel
// in the polygon rectangle bounds.
func (p Polygon) EachPixel(m draw.Image, fn func(x, y int)) {
	if len(p) < 3 {
		return
	}

	EachPixel(m, p.Rect().Bounds(), func(x, y int) {
		u := V(float64(x), float64(y))

		if u.In(p) {
			fn(x, y)
		}
	})
}

// Fill polygon on the image with the given color
func (p Polygon) Fill(m draw.Image, c color.Color) {
	p.EachPixel(m, func(x, y int) {
		Mix(m, x, y, c)
	})
}

// In returns true if the vector is inside the given polygon.
func (u Vec) In(p Polygon) bool {
	if len(p) < 3 {
		return false
	}

	a := p[0]

	in := rayIntersectsSegment(u, p[len(p)-1], a)

	for _, b := range p[1:] {
		if rayIntersectsSegment(u, a, b) {
			in = !in
		}

		a = b
	}

	return in
}

// Points are a list of points.
type Points []image.Point

// Polygon based on the points.
func (pts Points) Polygon() Polygon {
	var p Polygon

	for i := range pts {
		p = append(p, PV(pts[i]))
	}

	return p
}

// Segment intersect expression from
// https://www.ecse.rpi.edu/Homepages/wrf/Research/Short_Notes/pnpoly.html
//
// Currently the compiler inlines the function by default.
func rayIntersectsSegment(u, a, b Vec) bool {
	return (a.Y > u.Y) != (b.Y > u.Y) &&
		u.X < (b.X-a.X)*(u.Y-a.Y)/(b.Y-a.Y)+a.X
}
