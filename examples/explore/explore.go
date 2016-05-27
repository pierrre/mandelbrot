package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	mandelbrot_examples "github.com/pierrre/mandelbrot/examples"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
	mandelbrot_image_colorizer_rainbow "github.com/pierrre/mandelbrot/image/colorizer/rainbow"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	size := image.Pt(512, 512)
	rotate := 0.0
	baseScale := 1.0
	translate := complex(0, 0)
	steps, stepScale := 50, 2.0 // 50,2.0 | 85,1.5 | 155,1.25

	boundedColor := color.Black
	clr := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(boundedColor),
		mandelbrot_image_colorizer_rainbow.Colorizer(16, 0),
	)

	im := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	for step := 0; step < steps; step++ {
		scale := baseScale * mandelbrot_image.ImageScale(size) * math.Pow(stepScale, float64(step))
		tsf := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
		maxIter := mandelbrot_image.MaxIter(scale)

		log.Printf("step=%d translate=%g scale=%v maxIter=%d", step, translate, scale, maxIter)

		mandelbrot_image.RenderParallel(im, tsf, maxIter, clr)

		file := fmt.Sprintf("explore_%04d.png", step)
		mandelbrot_examples.Save(im, file)

		p := findBorderBoundedPoint(im, boundedColor)
		translate = tsf(complex(float64(p.X), float64(p.Y)))
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