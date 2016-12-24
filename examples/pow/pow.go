package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/disintegration/gift"
	"github.com/pierrre/mandelbrot"
	mandelbrot_examples "github.com/pierrre/mandelbrot/examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
	mandelbrot_image_colorizer_rainbow "github.com/pierrre/mandelbrot/image/colorizer/rainbow"
)

func main() {
	start := float64(2)
	stop := float64(3)
	steps := 1000

	size := image.Pt(2048, 2048)
	rotate := 0.0
	scale := 1.0
	translate := complex(0, 0)

	smooth := uint(2)
	smoothSize := size.Mul(1 << smooth)
	im := image.NewRGBA(image.Rect(0, 0, smoothSize.X, smoothSize.Y))
	scale *= mandelbrot_image.ImageScale(smoothSize)
	tsf := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale) * 10
	clr := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(color.Black),
		mandelbrot_image_colorizer_rainbow.Colorizer(16, 0),
	)

	for step := 0; step < steps; step++ {
		pow := start + ((stop - start) / float64(steps-1) * float64(step))
		log.Printf("%d: %f", step, pow)

		f := mandelbrot.NewPow(maxIter, pow)
		mandelbrot_image.RenderParallel(im, tsf, f, clr)

		tmp := im
		if smooth > 0 {
			g := gift.New(gift.Resize(size.X, size.Y, gift.LanczosResampling))
			tmp = image.NewRGBA(g.Bounds(im.Bounds()))
			g.Draw(tmp, im)
		}

		file := fmt.Sprintf("pow_%04d.png", step)
		mandelbrot_examples.Save(tmp, file)
	}

}
