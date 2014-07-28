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
				OK:   false,
				Iter: iter,
				Abs:  0,
			}
		}
	}
	return Result{
		OK:   true,
		Iter: iter,
		Abs:  abs,
	}
}

type Result struct {
	OK   bool
	Iter uint
	Abs  float64
}
