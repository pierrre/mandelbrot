package main

import (
	"context"
	"image"
	"image/color"

	"github.com/disintegration/gift"
	"github.com/pierrre/mandelbrot"
	mandelbrot_examples "github.com/pierrre/mandelbrot/examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
	mandelbrot_image_colorizer_rainbow "github.com/pierrre/mandelbrot/image/colorizer/rainbow"
)

func main() {
	ctx := context.Background()

	size := image.Pt(4096, 4096)
	rotate := 0.0
	scale := 1.6
	translate := complex(-0.75, 0)
	smooth := uint(1)

	smoothSize := size.Mul(1 << smooth)
	im := image.NewRGBA(image.Rect(0, 0, smoothSize.X, smoothSize.Y))

	scale *= mandelbrot_image.ImageScale(smoothSize)
	tsf := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	f := mandelbrot.New(maxIter)
	clr := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(color.Black),
		mandelbrot_image_colorizer_rainbow.Colorizer(16, 0),
	)
	mandelbrot_image.RenderParallel(ctx, im, tsf, f, clr)

	if smooth > 0 {
		g := gift.New(gift.Resize(size.X, size.Y, gift.LanczosResampling))
		tmp := image.NewRGBA(g.Bounds(im.Bounds()))
		g.Draw(tmp, im)
		im = tmp
	}

	mandelbrot_examples.Save(im, "color.png")
}
