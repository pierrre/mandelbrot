package mandelbrot

import "math"

func Mandelbrot(c complex128, maxIter int) Result {
	// optimization: skip first bulb/cardioid
	const quarter = 1.0 / 4.0
	foo := real(c) - quarter
	imagCSquare := imag(c) * imag(c)
	q := foo*foo + imagCSquare
	if q*(q+foo) < imagCSquare*quarter {
		return Result{
			Bounded: true,
			Iter:    maxIter,
			Abs:     0,
		}
	}

	z := c
	iter := 0
	var absSquare float64
	for {
		if iter >= maxIter {
			return Result{
				Bounded: true,
				Iter:    maxIter,
			}
		}

		// optimization: calculate "abs square" instead of "abs"
		absSquare = real(z)*real(z) + imag(z)*imag(z)
		if absSquare > 4 {
			return Result{
				Bounded: false,
				Iter:    iter,
				Abs:     math.Sqrt(absSquare),
			}
		}

		z = z*z + c
		iter++
	}
}

type Result struct {
	Bounded bool
	Iter    int
	Abs     float64
}
