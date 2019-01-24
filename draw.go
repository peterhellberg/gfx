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

// DrawOver draws src over dst.
func DrawOver(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	draw.Draw(dst, r, src, sp, draw.Over)
}

// DrawSrc draws src on dst.
func DrawSrc(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	draw.Draw(dst, r, src, sp, draw.Src)
}

// DrawOverPalettedImage draws a PalettedImage over a PalettedDrawImage.
func DrawOverPalettedImage(dst PalettedDrawImage, r image.Rectangle, src PalettedImage) {
	w, h, m := r.Dx(), r.Dy(), r.Min

	for x := m.X; x != w; x++ {
		for y := m.Y; y != h; y++ {
			if src.AlphaAt(x, y) > 0 {
				dst.SetColorIndex(x, y, src.ColorIndexAt(x, y))
			}
		}
	}
}

// DrawLayerOverPaletted draws a *Layer over a *Paletted.
// (slightly faster than using the generic DrawOverPalettedImage)
func DrawLayerOverPaletted(dst *Paletted, r image.Rectangle, src *Layer) {
	w, h, m := r.Dx(), r.Dy(), r.Min

	for x := m.X; x != w; x++ {
		for y := m.Y; y != h; y++ {
			if src.AlphaAt(x, y) > 0 {
				dst.SetColorIndex(x, y, src.ColorIndexAt(x, y))
			}
		}
	}
}

// DrawLine draws a line of the given color.
// A thickness of <= 1 is drawn using DrawBresenhamLine.
func DrawLine(dst draw.Image, from, to Vec, thickness float64, c color.Color) {
	if thickness <= 1 {
		DrawBresenhamLine(dst, from, to, c)
		return
	}

	polylineFromTo(from, to, thickness).Fill(dst, c)
}

// DrawImageRectangle draws a rectangle of the given color on the image.
func DrawImageRectangle(dst draw.Image, r image.Rectangle, c color.Color) {
	draw.Draw(dst, r, &image.Uniform{c}, image.ZP, draw.Over)
}

// DrawPolygon filled or as line polygons if the thickness is >= 1.
func DrawPolygon(dst draw.Image, p Polygon, thickness float64, c color.Color) {
	n := len(p)

	if n < 3 {
		return
	}

	switch {
	case thickness < 1:
		p.Fill(dst, c)
	default:
		for i := 0; i < n; i++ {
			if i+1 == n {
				polylineFromTo(p[n-1], p[0], thickness).Fill(dst, c)
			} else {
				polylineFromTo(p[i], p[i+1], thickness).Fill(dst, c)
			}
		}
	}
}

// DrawPolyline draws a polyline with the given color and thickness.
func DrawPolyline(dst draw.Image, pl Polyline, thickness float64, c color.Color) {
	for _, p := range pl {
		DrawPolygon(dst, p, thickness, c)
	}
}

// DrawCircle draws a circle with radius and thickness. (filled if thickness == 0)
func DrawCircle(dst draw.Image, u Vec, radius, thickness float64, c color.Color) {
	if thickness == 0 {
		DrawFilledCircle(dst, u, radius, c)
		return
	}

	bounds := IR(int(u.X-radius), int(u.Y-radius), int(u.X+radius), int(u.Y+radius))

	EachPixel(dst, bounds, func(x, y int) {
		v := V(float64(x), float64(y))

		l := u.To(v).Len() + 0.5

		if l < radius && l > radius-thickness {
			Mix(dst, x, y, c)
		}
	})
}

// DrawFilledCircle draws a filled circle.
func DrawFilledCircle(dst draw.Image, u Vec, radius float64, c color.Color) {
	bounds := IR(int(u.X-radius), int(u.Y-radius), int(u.X+radius), int(u.Y+radius))

	EachPixel(dst, bounds, func(x, y int) {
		v := V(float64(x), float64(y))

		if u.To(v).Len() < radius {
			Mix(dst, x, y, c)
		}
	})
}

// DrawFastFilledCircle draws a (crude) filled circle.
func DrawFastFilledCircle(dst draw.Image, u Vec, radius float64, c color.Color) {
	ir := int(radius)
	r2 := ir * ir
	pt := u.Pt()

	for y := -ir; y <= ir; y++ {
		for x := -ir; x <= ir; x++ {
			if x*x+y*y <= r2 {
				SetPoint(dst, pt.Add(Pt(x, y)), c)
			}
		}
	}
}

// DrawPointCircle draws a circle at the given point.
func DrawPointCircle(dst draw.Image, p image.Point, radius, thickness float64, c color.Color) {
	points := circlePoints(p, int(radius))

	switch {
	case thickness <= 1:
		for i := range points {
			SetPoint(dst, points[i], c)
		}
	default:
		center := PV(p)

		for i := range points {
			from := PV(points[i])
			to := from.Add(from.To(center).Unit().Scaled(thickness))

			DrawLine(dst, from, to, thickness, c)
		}
	}
}

func circlePoints(p image.Point, radius int) Points {
	var cp []image.Point

	x, y, dx, dy := radius-1, 0, 1, 1

	e := dx - (radius << 1)

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
			e += dx - (radius << 1)
		}
	}

	return cp
}
