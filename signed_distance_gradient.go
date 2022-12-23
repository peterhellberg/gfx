package gfx

type SignedDistanceGradient struct {
	Vec
}

func (sdg SignedDistanceGradient) Circle(r float64) Vec3 {
	d := sdg.Len()
	v := sdg.Div(d)

	return V3(d-r, v.X, v.Y)
}
