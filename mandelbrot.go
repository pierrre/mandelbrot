package mandelbrot

import (
	"math"
)

type Mandelbroter interface {
	Mandelbrot(c complex128) Result
}

type MandelbroterFunc func(c complex128) Result

func (f MandelbroterFunc) Mandelbrot(c complex128) Result {
	return f(c)
}

func Mandelbrot(maxIter int) Mandelbroter {
	return MandelbroterFunc(func(c complex128) Result {
		z := complex(0, 0)
		var iter int
		var absSquare float64
		for iter = 0; iter < maxIter; iter++ {
			z = z*z + c
			// cmplx.Abs(z) is 2x slower
			absSquare = real(z)*real(z) + imag(z)*imag(z)
			if absSquare > 4 { // math.Sqrt(4) == 2
				break
			}
		}
		return Result{
			Bounded: iter == maxIter,
			Iter:    iter + 1,
			Abs:     math.Sqrt(absSquare),
		}
	})
}

type Result struct {
	Bounded bool
	Iter    int
	Abs     float64
}
