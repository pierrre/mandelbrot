package mandelbrot

import (
	"strconv"
	"testing"
)

func BenchmarkNormal(b *testing.B) {
	f := New(1000)
	c := complex(-1, 0.15)
	for b.Loop() {
		f(c)
	}
}

func BenchmarkPow(b *testing.B) {
	for _, pow := range []float64{
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		13,
		14,
		15,
		16,
		17,
		18,
		19,
		20,
		21,
		100,
		2.5,
		21.5,
	} {
		b.Run(strconv.FormatFloat(pow, 'f', -1, 64), func(b *testing.B) {
			f := NewPow(1000, pow)
			c := complex(0.1, 0.1)
			for b.Loop() {
				f(c)
			}
		})
	}
}
