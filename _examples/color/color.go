package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/nfnt/resize"
	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	size := image.Pt(4096, 4096)
	rotate := 1.55
	scale := 2.0
	translate := complex(-0.5, 0)
	smooth := uint(1)

	smoothSize := size.Mul(1 << smooth)
	var im draw.Image = image.NewRGBA(image.Rect(0, 0, smoothSize.X, smoothSize.Y))

	scale *= mandelbrot_image.ImageScale(smoothSize)
	trans := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	colorizer := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(color.Black),
		mandelbrot_image.RainbowUnboundedColorizer(),
	)
	mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, colorizer)

	if smooth > 0 {
		im = resize.Resize(uint(size.X), uint(size.Y), im, resize.Lanczos3).(draw.Image)
	}

	mandelbrot_examples.Save(im, "color.png")
}
