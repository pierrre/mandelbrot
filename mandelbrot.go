package mandelbrot

import (
	"math/cmplx"
)

func Mandelbrot(c complex128, maxIter uint) Result {
	z := complex(0, 0)
	var iter uint
	var abs float64
	for iter = 0; iter < maxIter; iter++ {
		z = z*z + c
		abs = cmplx.Abs(z)
		if abs > 2 {
			return Result{
				Bounded: false,
				Iter:    iter,
				Abs:     0,
			}
		}
	}
	return Result{
		Bounded: true,
		Iter:    iter,
		Abs:     abs,
	}
}

type Result struct {
	Bounded bool
	Iter    uint
	Abs     float64
}
