// Package simple provides an example of rendering a simple Mandelbrot image.
package main

import (
	"image"

	"github.com/pierrre/mandelbrot"
	mandelbrot_cmd "github.com/pierrre/mandelbrot/cmd"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	size := image.Pt(1024, 1024)
	rotate := 0.0
	scale := 1.6
	translate := complex(-0.75, 0)

	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))

	scale *= mandelbrot_image.Scale(size)
	tsf := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	f := mandelbrot.New(maxIter)
	clr := mandelbrot_image.BWColorizer(false)
	mandelbrot_image.RenderParallel(im, tsf, f, clr)

	mandelbrot_cmd.Save(im, "simple.png")
}
