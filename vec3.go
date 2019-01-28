package gfx

import "math"

// Vec3 is a 3D vector type with X, Y and Z coordinates.
//
// Create vectors with the V3 constructor:
//
//   u := gfx.V3(1, 2, 3)
//   v := gfx.V3(8, -3, 4)
//
type Vec3 struct {
	X, Y, Z float64
}

// ZV3 is the zero Vec3
var ZV3 = Vec3{0, 0, 0}

// V3 is shorthand for Vec3{X: x, Y: y, Z: z}.
func V3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// Add returns the vector v+w.
func (u Vec3) Add(v Vec3) Vec3 {
	return Vec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

// Sub returns the vector v-w.
func (u Vec3) Sub(v Vec3) Vec3 {
	return Vec3{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

// Mul returns the vector v*s.
func (u Vec3) Mul(s float64) Vec3 {
	return Vec3{u.X * s, u.Y * s, u.Z * s}
}

// Div returns the vector v/s.
func (u Vec3) Div(s float64) Vec3 {
	return Vec3{u.X / s, u.Y / s, u.Z / s}
}

// Dot returns the dot (a.k.a. scalar) product of v and w.
func (u Vec3) Dot(v Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// CompMul returns the component-wise multiplication of two vectors.
func (u Vec3) CompMul(v Vec3) Vec3 {
	return Vec3{u.X * v.X, u.Y * v.Y, u.Z * v.Z}
}

// SqDist returns the square of the euclidian distance between two vectors.
func (u Vec3) SqDist(v Vec3) float64 {
	return u.Sub(v).SqLen()
}

// Dist returns the euclidian distance between two vectors.
func (u Vec3) Dist(v Vec3) float64 {
	return u.Sub(v).Len()
}

// SqLen returns the square of the length (euclidian norm) of a vector.
func (u Vec3) SqLen() float64 {
	return u.Dot(u)
}

// Len returns the length (euclidian norm) of a vector.
func (u Vec3) Len() float64 {
	return math.Sqrt(u.SqLen())
}

// Norm returns the normalized vector of a vector.
func (u Vec3) Norm() Vec3 {
	return u.Div(u.Len())
}
