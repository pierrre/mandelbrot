package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	width := 1024
	height := 1024
	im := image.NewGray(image.Rect(0, 0, width, height))

	scale := mandelbrot_image.ImageScale(im)
	scale *= 2
	translate := complex(-0.5, 0)
	trans := mandelbrot_image.BaseTransformation(im, scale, translate)

	maxIter := mandelbrot_image.MaxIter(scale)

	mandelbrot_image.RenderWorkerAuto(im, trans, maxIter, mandelbrot_image.BWColorizer)

	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("simple.png", buf.Bytes(), os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
