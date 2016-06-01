package image

import (
	"context"
	"image"
	"testing"

	"github.com/pierrre/mandelbrot"
)

func BenchmarkRender(b *testing.B) {
	ctx := context.Background()
	size := image.Pt(256, 256)
	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	rotate := 1.0
	scale := ImageScale(size)
	translate := complex(0, 0)
	tsf := BaseTransformation(im, rotate, scale, translate)
	maxIter := 500
	f := mandelbrot.New(maxIter)
	clr := BWColorizer(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Render(ctx, im, tsf, f, clr)
	}
}

func BenchmarkRenderParallel(b *testing.B) {
	ctx := context.Background()
	size := image.Pt(512, 512)
	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))
	rotate := 1.0
	scale := ImageScale(size)
	translate := complex(0, 0)
	tsf := BaseTransformation(im, rotate, scale, translate)
	maxIter := 500
	f := mandelbrot.New(maxIter)
	clr := BWColorizer(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RenderParallel(ctx, im, tsf, f, clr)
	}
}
