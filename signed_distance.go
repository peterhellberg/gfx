package gfx

import "math"

// SignedDistance holds 2D signed distance functions based on
// https://iquilezles.org/www/articles/distfunctions2d/distfunctions2d.htm
type SignedDistance struct {
	Vec
}

// Circle primitive
func (sd SignedDistance) Circle(r float64) float64 {
	return sd.Len() - r
}

// Line primitive
func (sd SignedDistance) Line(a, b Vec) float64 {
	pa, ba := sd.Sub(a), b.Sub(a)
	c := Clamp(pa.Dot(ba)/ba.Dot(ba), 0.0, 1.0)

	return pa.Sub(ba.Scaled(c)).Len()
}

// Rectangle primitive
func (sd SignedDistance) Rectangle(b Vec) float64 {
	d := sd.Abs().Sub(b)

	return d.Max(ZV).Len() + MathMin(MathMax(d.X, d.Y), 0)
}

// Rhombus primitive
func (sd SignedDistance) Rhombus(b Vec) float64 {
	q := sd.Abs()
	x := (-2*q.Normal().Dot(b.Normal()) + b.Normal().Dot(b.Normal())) / b.Dot(b)
	h := Clamp(x, -1.0, 1.0)
	d := q.Sub(b.Scaled(0.5).ScaledXY(V(1.0-h, 1.0+h))).Len()

	return d * Sign(q.X*b.Y+q.Y*b.X-b.X*b.Y)
}

// Rounded signed distance function shape
func (sd SignedDistance) Rounded(v, r float64) float64 {
	return v - r
}

// Annular signed distance function shape
func (sd SignedDistance) Annular(v, r float64) float64 {
	return math.Abs(v) - r
}
