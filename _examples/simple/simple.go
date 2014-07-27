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
	maxIter := 20

	size := 1024
	bounds := image.Rect(0, 0, size, size)
	im := image.NewGray(bounds)

	proj := mandelbrot_image.ProjectionFunc(func(x, y int) complex128 {
		return complex(
			(float64(x)/float64(bounds.Dx())*4)-2,
			-((float64(y) / float64(bounds.Dy()) * 4) - 2),
		)
	})

	mandelbrot_image.RenderWorkerAuto(im, proj, maxIter)

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
