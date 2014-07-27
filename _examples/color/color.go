package main

import (
	"image"
	"image/color"
	"math"

	"github.com/pierrre/mandelbrot"
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

	colorSin := func(x float64) uint8 {
		x = math.Sin(x)
		x = (x + 1) / 2
		return uint8(x * 255)
	}
	colorRainbow := func(x float64) color.Color {
		return color.RGBA{
			R: colorSin(x + (math.Pi * 0 / 3)),
			G: colorSin(x + (math.Pi * 2 / 3)),
			B: colorSin(x + (math.Pi * 4 / 3)),
			A: 255,
		}
	}
	colorizer := mandelbrot_image.ColorizerFunc(func(res mandelbrot.Result) color.Color {
		if res.OK {
			return color.Black
		}
		return colorRainbow(float64(res.Iter) / 4)
	})

	mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, colorizer)

	mandelbrot_examples.Save(im, "color.png")
}
