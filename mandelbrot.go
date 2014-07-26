package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/cmplx"
	"os"
)

func main() {
	maxIter := 50

	size := 1024
	bounds := image.Rect(0, 0, size, size)
	im := image.NewRGBA(bounds)

	proj := ProjectionFunc(func(x, y int) complex128 {
		return complex(
			(float64(x)/float64(bounds.Dx())*4)-2,
			-((float64(y) / float64(bounds.Dy()) * 4) - 2),
		)
	})

	MandelbrotImage(im, proj, maxIter)

	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("mandelbrot.png", buf.Bytes(), os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}

func MandelbrotImage(im draw.Image, proj Projection, maxIter int) {
	bounds := im.Bounds()
	boundsDy := bounds.Dy()
	boundsDx := bounds.Dx()
	for y := 0; y < boundsDy; y++ {
		for x := 0; x < boundsDx; x++ {
			c := proj.Project(x, y)
			var col color.Color
			if Mandelbrot(c, maxIter) {
				col = color.White
			} else {
				col = color.Black
			}
			im.Set(x, y, col)
		}
	}
}

func Mandelbrot(c complex128, maxIter int) bool {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return false
		}
	}
	return true
}

type Projection interface {
	Project(x, y int) complex128
}

type ProjectionFunc func(x, y int) complex128

func (pf ProjectionFunc) Project(x, y int) complex128 {
	return pf(x, y)
}
