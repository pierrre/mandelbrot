package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/nfnt/resize"
	"github.com/pierrre/mandelbrot"
	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
	mandelbrot_image_colorizer_rainbow "github.com/pierrre/mandelbrot/image/colorizer/rainbow"
)

func main() {
	size := image.Pt(4096, 4096)
	rotate := 0.0
	scale := 1.6
	translate := complex(-0.75, 0)
	smooth := uint(1)

	smoothSize := size.Mul(1 << smooth)
	var im draw.Image = image.NewRGBA(image.Rect(0, 0, smoothSize.X, smoothSize.Y))

	scale *= mandelbrot_image.ImageScale(smoothSize)
	trans := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	mandelbroter := mandelbrot.Mandelbrot(maxIter)
	colorizer := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(color.Black),
		mandelbrot_image_colorizer_rainbow.RainbowIterColorizer(16, 0),
	)
	mandelbrot_image.RenderWorkerAuto(mandelbroter, im, trans, colorizer)

	if smooth > 0 {
		im = resize.Resize(uint(size.X), uint(size.Y), im, resize.Lanczos3).(draw.Image)
	}

	mandelbrot_examples.Save(im, "color.png")
}
