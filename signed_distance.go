package gfx

// SignedDistance holds 2D signed distance functions based on
// https://iquilezles.org/www/articles/distfunctions2d/distfunctions2d.htm
type SignedDistance struct {
	Vec
}

// SignedDistanceFunc is a func that takes a SignedDistance and returns a float64.
type SignedDistanceFunc func(SignedDistance) float64

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

// EquilateralTriangle primitive
func (sd SignedDistance) EquilateralTriangle(s float64) float64 {
	k := MathSqrt(3)

	p := sd.Vec

	p.X = MathAbs(p.X) - s
	p.Y = p.Y + s/k

	if p.X+k*p.Y > 0.0 {
		p = V(p.X-k*p.Y, -k*p.X-p.Y).Scaled(0.5)
	}

	p.X -= Clamp(p.X, -2.0, 0.0)

	return -p.Len() * Sign(p.Y)
}

// IsoscelesTriangle primitive
func (sd SignedDistance) IsoscelesTriangle(q Vec) float64 {
	p := sd.Vec

	p.X = MathAbs(p.X)

	a := p.Sub(q.Scaled(Clamp(p.Dot(q)/q.Dot(q), 0.0, 1.0)))
	b := p.Sub(q.ScaledXY(V(Clamp(p.X/q.X, 0.0, 1.0), 1.0)))

	s := -Sign(q.Y)

	d := V(a.Dot(a), s*(p.X*q.Y-p.Y*q.X)).Min(V(b.Dot(b), s*(p.Y-q.Y)))

	return -MathSqrt(d.X) * Sign(d.Y)
}

// Rounded signed distance function shape
func (sd SignedDistance) Rounded(v, r float64) float64 {
	return v - r
}

// Annular signed distance function shape
func (sd SignedDistance) Annular(v, r float64) float64 {
	return MathAbs(v) - r
}

// OpUnion basic boolean operation for union.
func (sd SignedDistance) OpUnion(x, y float64) float64 {
	return MathMin(x, y)
}

// OpSubtraction basic boolean operation for subtraction.
func (sd SignedDistance) OpSubtraction(x, y float64) float64 {
	return MathMax(-x, y)
}

// OpIntersection basic boolean operation for intersection.
func (sd SignedDistance) OpIntersection(x, y float64) float64 {
	return MathMax(x, y)
}

// OpSmoothUnion smooth operation for union.
func (sd SignedDistance) OpSmoothUnion(x, y, k float64) float64 {
	h := Clamp(0.5+0.5*(y-x)/k, 0.0, 1.0)

	return Lerp(y, x, h) - k*h*(1.0-h)
}

// OpSmoothSubtraction smooth operation for subtraction.
func (sd SignedDistance) OpSmoothSubtraction(x, y, k float64) float64 {
	h := Clamp(0.5-0.5*(y+x)/k, 0.0, 1.0)

	return Lerp(y, -x, h) + k*h*(1.0-h)
}

// OpSmoothIntersection smooth operation for intersection.
func (sd SignedDistance) OpSmoothIntersection(x, y, k float64) float64 {
	h := Clamp(0.5-0.5*(y-x)/k, 0.0, 1.0)

	return Lerp(y, x, h) + k*h*(1.0-h)
}

// OpSymX symmetry operation for X.
func (sd SignedDistance) OpSymX(sdf SignedDistanceFunc) float64 {
	sd.X = MathAbs(sd.X)

	return sdf(sd)
}

// OpSymY symmetry operation for Y.
func (sd SignedDistance) OpSymY(sdf SignedDistanceFunc) float64 {
	sd.Y = MathAbs(sd.Y)

	return sdf(sd)
}

// OpSymXY symmetry operation for X and Y.
func (sd SignedDistance) OpSymXY(sdf SignedDistanceFunc) float64 {
	sd.X = MathAbs(sd.X)
	sd.Y = MathAbs(sd.Y)

	return sdf(sd)
}

// OpRepeat repeats based on the given c vector.
func (sd SignedDistance) OpRepeat(c Vec, sdf SignedDistanceFunc) float64 {
	q := sd.Mod(c).Sub(c.Scaled(0.5))

	return sdf(SignedDistance{q})
}
