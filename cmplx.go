package gfx

import "math/cmplx"

// CmplxSin returns the sine of x.
func CmplxSin(x complex128) complex128 {
	return cmplx.Sin(x)
}

// CmplxSinh returns the hyperbolic sine of x.
func CmplxSinh(x complex128) complex128 {
	return cmplx.Sinh(x)
}

// CmplxCos returns the cosine of x.
func CmplxCos(x complex128) complex128 {
	return cmplx.Cos(x)
}

// CmplxCosh returns the hyperbolic cosine of x.
func CmplxCosh(x complex128) complex128 {
	return cmplx.Cosh(x)
}

// CmplxPow returns x**y, the base-x exponential of y.
func CmplxPow(x, y complex128) complex128 {
	return cmplx.Pow(x, y)
}

// CmplxPhase returns the phase (also called the argument) of x.
// The returned value is in the range [-Pi, Pi].
func CmplxPhase(x complex128) float64 {
	return cmplx.Phase(x)
}
