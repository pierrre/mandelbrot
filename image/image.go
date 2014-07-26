package image

import (
	"image/color"
	"image/draw"

	"github.com/pierrre/mandelbrot"
)

func Render(im draw.Image, proj Projection, maxIter int) {
	bounds := im.Bounds()
	boundsDy := bounds.Dy()
	boundsDx := bounds.Dx()
	for y := 0; y < boundsDy; y++ {
		for x := 0; x < boundsDx; x++ {
			c := proj.Project(x, y)
			var col color.Color
			if mandelbrot.Mandelbrot(c, maxIter) {
				col = color.White
			} else {
				col = color.Black
			}
			im.Set(x, y, col)
		}
	}
}

type Projection interface {
	Project(x, y int) complex128
}

type ProjectionFunc func(x, y int) complex128

func (pf ProjectionFunc) Project(x, y int) complex128 {
	return pf(x, y)
}
