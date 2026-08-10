// Package mandelbrot provides utilities to compute the Mandelbrot set.
package mandelbrot

import (
	"math"
	"math/cmplx"
)

// Func represents a function that computes the Mandelbrot set for a given point.
type Func func(complex128) Result

// Result represents a result of the Mandelbrot set computation for a point.
type Result struct {
	Bounded bool
	Iter    int
	Abs     float64
}

// New returns a new [Func].
func New(maxIter int) Func {
	return newPow2(maxIter)
}

// NewPow returns a new [Func] that uses the given power.
func NewPow(maxIter int, pow float64) Func {
	if f, ok := newPows[pow]; ok {
		return f(maxIter)
	}
	return newPow(maxIter, pow)
}

var newPows = map[float64]func(maxIter int) Func{
	2:  newPow2,
	3:  newPow3,
	4:  newPow4,
	5:  newPow5,
	6:  newPow6,
	7:  newPow7,
	8:  newPow8,
	9:  newPow9,
	10: newPow10,
	11: newPow11,
	12: newPow12,
	13: newPow13,
	14: newPow14,
	15: newPow15,
	16: newPow16,
	17: newPow17,
	18: newPow18,
	19: newPow19,
	20: newPow20,
}

func newPow2(maxIter int) Func {
	return func(c complex128) Result {
		// optimization: skip first bulb/cardioid
		const quarter = 1.0 / 4.0
		foo := real(c) - quarter
		imagCSquare := imag(c) * imag(c)
		q := foo*foo + imagCSquare
		if q*(q+foo) < imagCSquare*quarter {
			return Result{
				Bounded: true,
				Iter:    maxIter,
			}
		}

		z := c
		for iter := range maxIter {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z = z*z + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPowFunc(maxIter int, pow func(complex128) complex128) Func {
	return func(c complex128) Result {
		z := c
		for iter := range maxIter {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z = pow(z) + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow3(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		return z * z2
	})
}

func newPow4(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		return z2 * z2
	})
}

func newPow5(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		return z * z4
	})
}

func newPow6(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		return z2 * z4
	})
}

func newPow7(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z6 := z2 * z4
		return z * z6
	})
}

func newPow8(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		return z4 * z4
	})
}

func newPow9(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		return z * z8
	})
}

func newPow10(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		return z2 * z8
	})
}

func newPow11(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z10 := z2 * z8
		return z * z10
	})
}

func newPow12(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		return z4 * z8
	})
}

func newPow13(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z12 := z4 * z8
		return z * z12
	})
}

func newPow14(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z12 := z4 * z8
		return z2 * z12
	})
}

func newPow15(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z5 := z * z4
		z10 := z5 * z5
		return z5 * z10
	})
}

func newPow16(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		return z8 * z8
	})
}

func newPow17(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z16 := z8 * z8
		return z * z16
	})
}

func newPow18(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z16 := z8 * z8
		return z2 * z16
	})
}

func newPow19(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z16 := z8 * z8
		z18 := z2 * z16
		return z * z18
	})
}

func newPow20(maxIter int) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		z2 := z * z
		z4 := z2 * z2
		z8 := z4 * z4
		z16 := z8 * z8
		return z4 * z16
	})
}

func newPow(maxIter int, pow float64) Func {
	return newPowFunc(maxIter, func(z complex128) complex128 {
		return cmplx.Pow(z, complex(pow, 0))
	})
}
