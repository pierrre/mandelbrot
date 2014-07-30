package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"

	mandelbrot_examples "github.com/pierrre/mandelbrot/_examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	size := image.Pt(1024, 1024)
	scale := 1.0
	translate := complex(0, 0)

	boundedColor := color.Black
	colorizer := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(boundedColor),
		mandelbrot_image.RainbowUnboundedColorizer(),
	)

	scale *= mandelbrot_image.ImageScale(size)

	for i := 0; i < 50; i++ {
		var im draw.Image = image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
		trans := mandelbrot_image.BaseTransformation(im, scale, translate)
		maxIter := mandelbrot_image.MaxIter(scale)
		fmt.Println(i, translate, scale)
		mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, colorizer)

		mandelbrot_examples.Save(im, fmt.Sprintf("explore_%04d.png", i))

		p := findBorderBoundedPoint(im, boundedColor)
		translate = trans.Transform(complex(float64(p.X), float64(p.Y)))
		scale *= 2.0
	}
}

func findBorderBoundedPoint(im image.Image, boundedColor color.Color) image.Point {
	for {
		p := findBoundedPoint(im, boundedColor)
		if checkBorderPoint(im, boundedColor, p) {
			return p
		}
	}
}

func findBoundedPoint(im image.Image, boundedColor color.Color) image.Point {
	size := im.Bounds().Size()
	for {
		x := rand.Intn(size.X)
		y := rand.Intn(size.Y)
		if colorEqualRGBA(im.At(x, y), boundedColor) {
			return image.Pt(x, y)
		}
	}
}

var checkBorderPointTrans = []image.Point{
	image.Pt(-1, -1),
	image.Pt(0, -1),
	image.Pt(1, -1),
	image.Pt(-1, 0),
	image.Pt(1, 0),
	image.Pt(-1, 1),
	image.Pt(0, 1),
	image.Pt(1, 1),
}

func checkBorderPoint(im image.Image, boundedColor color.Color, p image.Point) bool {
	bounds := im.Bounds()
	for _, trans := range checkBorderPointTrans {
		pTrans := p.Add(trans)
		if !pTrans.In(bounds) {
			continue
		}
		if !colorEqualRGBA(im.At(pTrans.X, pTrans.Y), boundedColor) {
			return true
		}
	}
	return false
}

func colorEqualRGBA(c1, c2 color.Color) bool {
	c1r, c1g, c1b, c1a := c1.RGBA()
	c2r, c2g, c2b, c2a := c2.RGBA()
	return c1r == c2r && c1g == c2g && c1b == c2b && c1a == c2a
}
