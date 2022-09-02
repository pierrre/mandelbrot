package image

import (
	"image"
	"image/draw"

	"github.com/pierrre/imageutil"
	"github.com/pierrre/mandelbrot"
)

// Render renders to an image.
func Render(im draw.Image, tsf Transformation, f mandelbrot.Func, clr Colorizer) {
	render(imageutil.NewSetFunc(im), im.Bounds(), tsf, f, clr)
}

// RenderParallel renders to an image in parallel.
func RenderParallel(im draw.Image, tsf Transformation, f mandelbrot.Func, clr Colorizer) {
	set := imageutil.NewSetFunc(im)
	imageutil.Parallel2D(im.Bounds(), func(bds image.Rectangle) {
		render(set, bds, tsf, f, clr)
	})
}

func render(set imageutil.SetFunc, bds image.Rectangle, tsf Transformation, f mandelbrot.Func, clr Colorizer) {
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
