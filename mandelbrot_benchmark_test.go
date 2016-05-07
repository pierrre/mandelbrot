package mandelbrot

import (
	"testing"
)

var benchRes Result

func BenchmarkMandelbrot(b *testing.B) {
	c := complex(0.1, 0.1)
	var res Result
	for i := 0; i < b.N; i++ {
		res = Mandelbrot(c, 100)
	}
	benchRes = res
}
