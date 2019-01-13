package gfx

import (
	"image"
	"image/color"
	"image/draw"
)

// Drawer is the draw interface.
type Drawer interface {
	Draw(draw.Image)
}

// DrawLine draws a line of the given color.
// A thickness of <= 1 is drawn using DrawBresenhamLine.
func DrawLine(m draw.Image, from, to Vec, t float64, c color.Color) {
	if t <= 1 {
		DrawBresenhamLine(m, from, to, c)
		return
	}

	polylineFromTo(from, to, t).Fill(m, c)
}

// DrawCircle draws a circle with radius and thickness.
func DrawCircle(m draw.Image, u Vec, c color.Color, r, t float64) {
	if t == 0 {
		DrawFilledCircle(m, u, c, r)
		return
	}

	bounds := IR(int(u.X-r), int(u.Y-r), int(u.X+r), int(u.Y+r))

	EachPixel(m, bounds, func(x, y int) {
		v := V(float64(x), float64(y))

		l := u.To(v).Len() + 0.5

		if l < r && l > r-t {
			Mix(m, x, y, c)
		}
	})
}

// DrawFilledCircle draws a filled circle.
func DrawFilledCircle(m draw.Image, u Vec, c color.Color, r float64) {
	bounds := IR(int(u.X-r), int(u.Y-r), int(u.X+r), int(u.Y+r))

	EachPixel(m, bounds, func(x, y int) {
		v := V(float64(x), float64(y))

		if u.To(v).Len() < r {
			Mix(m, x, y, c)
		}
	})
}

// DrawFastFilledCircle draws a (crude) filled circle.
func DrawFastFilledCircle(m draw.Image, u Vec, c color.Color, r float64) {
	ir := int(r)
	r2 := ir * ir
	pt := u.Pt()

	for y := -ir; y <= ir; y++ {
		for x := -ir; x <= ir; x++ {
			if x*x+y*y <= r2 {
				SetPoint(m, pt.Add(Pt(x, y)), c)
			}
		}
	}
}

// DrawPointCircle draws a circle at the given point
func DrawPointCircle(m draw.Image, p image.Point, c color.Color, r int, t float64) {
	points := circlePoints(p, r)

	switch {
	case t <= 1:
		for i := range points {
			SetPoint(m, points[i], c)
		}
	default:
		center := PV(p)

		for i := range points {
			from := PV(points[i])
			to := from.Add(from.To(center).Unit().Scaled(t))

			DrawLine(m, from, to, t, c)
		}
	}
}

func circlePoints(p image.Point, r int) Points {
	var cp []image.Point

	x, y, dx, dy := r-1, 0, 1, 1

	e := dx - (r << 1)

	for x >= y {
		cp = append(cp,
			p.Add(Pt(x, y)),
			p.Add(Pt(y, x)),
			p.Add(Pt(-y, x)),
			p.Add(Pt(-x, y)),
			p.Add(Pt(-x, -y)),
			p.Add(Pt(-y, -x)),
			p.Add(Pt(y, -x)),
			p.Add(Pt(x, -y)),
		)

		if e <= 0 {
			y++
			e += dy
			dy += 2
		}

		if e > 0 {
			x--
			dx += 2
			e += dx - (r << 1)
		}
	}

	return cp
}

// DrawImageRectangle draws a rectangle of the given color on the image.
func DrawImageRectangle(m draw.Image, r image.Rectangle, c color.Color) {
	draw.Draw(m, r, &image.Uniform{c}, image.ZP, draw.Over)
}

// DrawPolygon filled or as line polygons if the thickness is >= 1.
func DrawPolygon(m draw.Image, p Polygon, c color.Color, t float64) {
	n := len(p)

	if n < 3 {
		return
	}

	switch {
	case t < 1:
		p.Fill(m, c)
	default:
		for i := 0; i < n; i++ {
			if i+1 == n {
				polylineFromTo(p[n-1], p[0], t).Fill(m, c)
			} else {
				polylineFromTo(p[i], p[i+1], t).Fill(m, c)
			}
		}
	}
}

// DrawPolyline draws a polyline with the given color and thickness.
func DrawPolyline(m draw.Image, pl Polyline, c color.Color, t float64) {
	for _, p := range pl {
		DrawPolygon(m, p, c, t)
	}
}
