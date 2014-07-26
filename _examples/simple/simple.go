package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/nfnt/resize"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	maxIter := 100

	size := 16384
	bounds := image.Rect(0, 0, size, size)
	im := image.NewGray(bounds)

	proj := mandelbrot_image.ProjectionFunc(func(x, y int) complex128 {
		return complex(
			(float64(x)/float64(bounds.Dx())*4)-2,
			-((float64(y) / float64(bounds.Dy()) * 4) - 2),
		)
	})

	mandelbrot_image.Render(im, proj, maxIter)

	resizeSize := uint(8192)
	imResized := resize.Resize(resizeSize, resizeSize, im, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err := png.Encode(buf, imResized)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("simple.png", buf.Bytes(), os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
