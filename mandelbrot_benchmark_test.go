package mandelbrot

import (
	"testing"
)

func BenchmarkMandelbrot(b *testing.B) {
	m := NewMandelbroter(100)
	c := complex(0.1, 0.1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Mandelbrot(c)
		}
	})
}
