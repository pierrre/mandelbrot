package image

import (
	"image/color"

	"github.com/pierrre/mandelbrot"
)

type Colorizer interface {
	Colorize(mandelbrot.Result) color.Color
}

type ColorizerFunc func(mandelbrot.Result) color.Color

func (f ColorizerFunc) Colorize(res mandelbrot.Result) color.Color {
	return f(res)
}

var BWColorizer = ColorizerFunc(func(res mandelbrot.Result) color.Color {
	if res.OK {
		return color.White
	} else {
		return color.Black
	}
})
