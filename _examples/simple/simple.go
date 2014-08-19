package main

import (
	"image"

	"github.com/pierrre/mandelbrot"
	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	size := image.Pt(1024, 1024)
	rotate := 0.0
	scale := 1.6
	translate := complex(-0.75, 0)

	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))

	scale *= mandelbrot_image.ImageScale(size)
	trans := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	mandelbroter := mandelbrot.Mandelbrot(maxIter)
	colorizer := mandelbrot_image.BWColorizer(false)
	mandelbrot_image.RenderWorkerAuto(mandelbroter, im, trans, colorizer)

	mandelbrot_examples.Save(im, "simple.png")
}
