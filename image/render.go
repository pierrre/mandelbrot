package image

import (
	"context"
	"image"
	"image/draw"

	"github.com/pierrre/imageutil"
	"github.com/pierrre/mandelbrot"
)

func Render(ctx context.Context, im draw.Image, tsf Transformation, f mandelbrot.Func, clr Colorizer) {
	render(ctx, imageutil.NewSetFunc(im), im.Bounds(), tsf, f, clr)
}

func RenderParallel(ctx context.Context, im draw.Image, tsf Transformation, f mandelbrot.Func, clr Colorizer) {
	set := imageutil.NewSetFunc(im)
	imageutil.Parallel2D(ctx, im.Bounds(), func(ctx context.Context, bds image.Rectangle) {
		render(ctx, set, bds, tsf, f, clr)
	})
}

func render(ctx context.Context, set imageutil.SetFunc, bds image.Rectangle, tsf Transformation, f mandelbrot.Func, clr Colorizer) {
	for y := bds.Min.Y; y < bds.Max.Y; y++ {
		for x := bds.Min.X; x < bds.Max.X; x++ {
			c := complex(float64(x), float64(y))
			c = tsf(c)
			res := f(c)
			col := clr(c, res)
			r, g, b, a := col.RGBA()
			set(x, y, r, g, b, a)
		}
	}
}
