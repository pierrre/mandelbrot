package mandelbrot

import (
	"math/cmplx"
)

func Mandelbrot(c complex128, maxIter int) bool {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return false
		}
	}
	return true
}
