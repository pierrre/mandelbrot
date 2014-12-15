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

func NewMandelbroter(maxIter int) Mandelbroter {
	return MandelbroterFunc(func(c complex128) Result {
		return mandelbrot(c, maxIter)
	})
}

func mandelbrot(c complex128, maxIter int) Result {
	z := c
	var iter int = 0
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
