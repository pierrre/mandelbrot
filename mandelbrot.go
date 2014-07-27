package mandelbrot

import (
	"math/cmplx"
)

func Mandelbrot(c complex128, maxIter int) (ok bool, iter int, abs float64) {
	z := complex(0, 0)
	for iter = 0; iter < maxIter; iter++ {
		z = z*z + c
		abs = cmplx.Abs(z)
		if abs > 2 {
			return false, iter, 0
		}
	}
	return true, iter, abs
}
