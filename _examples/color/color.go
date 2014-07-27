package main

import (
	"image"
	"image/draw"

	"github.com/nfnt/resize"
	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	width := 8192
	height := 8192
	scale := 2.0
	translate := complex(-0.5, 0)
	smooth := uint(1)

	smoothMultiplicator := 1 << smooth
	smoothWidth := width * smoothMultiplicator
	smoothHeight := width * smoothMultiplicator
	var im draw.Image = image.NewRGBA(image.Rect(0, 0, smoothWidth, smoothHeight))

	scale *= mandelbrot_image.ImageScale(im)
	trans := mandelbrot_image.BaseTransformation(im, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	colorizer := mandelbrot_image.RainbowColorizer()
	mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, colorizer)

	if smooth > 0 {
		im = resize.Resize(uint(width), uint(height), im, resize.Lanczos3).(draw.Image)
	}

	mandelbrot_examples.Save(im, "color.png")
}
