package main

import (
	"image"

	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	width := 1024
	height := 1024
	scale := 2.0
	translate := complex(-0.5, 0)

	im := image.NewGray(image.Rect(0, 0, width, height))

	scale *= mandelbrot_image.ImageScale(im)
	trans := mandelbrot_image.BaseTransformation(im, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	colorizer := mandelbrot_image.BWColorizer()
	mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, colorizer)

	mandelbrot_examples.Save(im, "simple.png")
}
