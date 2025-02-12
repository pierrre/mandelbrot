package image

import (
	"image"
	"testing"

	"github.com/pierrre/mandelbrot"
)

func BenchmarkRender(b *testing.B) {
	size := image.Pt(256, 256)
	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	rotate := 1.0
	scale := Scale(size)
	translate := complex(0, 0)
	tsf := BaseTransformation(im, rotate, scale, translate)
	maxIter := 500
	f := mandelbrot.New(maxIter)
	clr := BWColorizer(false)
	for b.Loop() {
		Render(im, tsf, f, clr)
	}
}

func BenchmarkRenderParallel(b *testing.B) {
	size := image.Pt(512, 512)
	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	rotate := 1.0
	scale := Scale(size)
	translate := complex(0, 0)
	tsf := BaseTransformation(im, rotate, scale, translate)
	maxIter := 500
	f := mandelbrot.New(maxIter)
	clr := BWColorizer(false)
	for b.Loop() {
		RenderParallel(im, tsf, f, clr)
	}
}
