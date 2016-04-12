package mandelbrot

import (
	"testing"
)

func BenchmarkMandelbrot(b *testing.B) {
	m := NewMandelbroter(100)
	c := complex(0.1, 0.1)
	for i := 0; i < b.N; i++ {
		m.Mandelbrot(c)
	}
}
