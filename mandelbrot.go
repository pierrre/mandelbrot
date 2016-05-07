package mandelbrot

import (
	"math"
)

func Mandelbrot(c complex128, maxIter int) Result {
	z := c
	iter := 0
	var absSquare float64
	for {
		if iter >= maxIter {
			break
		}
		absSquare = real(z)*real(z) + imag(z)*imag(z)
		if absSquare > 4 {
			break
		}
		z = z*z + c
		iter++
	}
	return Result{
		Bounded: iter == maxIter,
		Iter:    iter,
		Abs:     math.Sqrt(absSquare),
	}
}

type Result struct {
	Bounded bool
	Iter    int
	Abs     float64
}
