package main

import (
	"image"

	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	width := 8192
	height := 8192
	im := image.NewRGBA(image.Rect(0, 0, width, height))

	scale := mandelbrot_image.ImageScale(im)
	scale *= 2
	translate := complex(-0.5, 0)
	trans := mandelbrot_image.BaseTransformation(im, scale, translate)

	maxIter := mandelbrot_image.MaxIter(scale)

	colorizer := mandelbrot_image.RainbowColorizer()

	mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, colorizer)

	mandelbrot_examples.Save(im, "color.png")
}
