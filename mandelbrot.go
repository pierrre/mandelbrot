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

// New returns a new Func.
func New(maxIter int) Func {
	return newPow2(maxIter)
}

// NewPow returns a new Func that uses the given power.
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
		for iter := 0; iter < maxIter; iter++ {
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

func newPow3(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z3 := z * z2
			z = z3 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow4(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z = z4 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow5(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z5 := z * z4
			z = z5 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow6(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z6 := z2 * z4
			z = z6 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow7(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z6 := z2 * z4
			z7 := z * z6
			z = z7 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow8(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z = z8 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow9(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z9 := z * z8
			z = z9 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow10(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z10 := z2 * z8
			z = z10 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow11(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z10 := z2 * z8
			z11 := z * z10
			z = z11 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow12(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z12 := z4 * z8
			z = z12 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow13(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z12 := z4 * z8
			z13 := z * z12
			z = z13 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow14(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z12 := z4 * z8
			z14 := z2 * z12
			z = z14 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow15(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z5 := z * z4
			z10 := z5 * z5
			z15 := z5 * z10
			z = z15 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow16(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z16 := z8 * z8
			z = z16 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow17(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z16 := z8 * z8
			z17 := z * z16
			z = z17 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow18(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z16 := z8 * z8
			z18 := z2 * z16
			z = z18 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow19(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z16 := z8 * z8
			z18 := z2 * z16
			z19 := z * z18
			z = z19 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow20(maxIter int) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z2 := z * z
			z4 := z2 * z2
			z8 := z4 * z4
			z16 := z8 * z8
			z20 := z4 * z16
			z = z20 + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}

func newPow(maxIter int, pow float64) Func {
	return func(c complex128) Result {
		z := c
		for iter := 0; iter < maxIter; iter++ {
			// optimization: calculate "abs square" instead of "abs"
			absSquare := real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 {
				return Result{
					Bounded: false,
					Iter:    iter,
					Abs:     math.Sqrt(absSquare),
				}
			}
			z = cmplx.Pow(z, complex(pow, 0)) + c
		}
		return Result{
			Bounded: true,
			Iter:    maxIter,
		}
	}
}
