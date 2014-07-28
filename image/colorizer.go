package image

import (
	"image/color"
	"math"

	"github.com/pierrre/mandelbrot"
)

type Colorizer interface {
	Colorize(mandelbrot.Result) color.Color
}

type ColorizerFunc func(mandelbrot.Result) color.Color

func (f ColorizerFunc) Colorize(res mandelbrot.Result) color.Color {
	return f(res)
}

func BWColorizer(invert bool) Colorizer {
	return ColorizerFunc(func(res mandelbrot.Result) color.Color {
		if res.Bounded != invert {
			return color.White
		} else {
			return color.Black
		}
	})
}

func RainbowColorizer() Colorizer {
	colorSin := func(x float64) uint8 {
		x = math.Sin(x)
		x = (x + 1) / 2
		return uint8(x * 255)
	}
	colorRainbow := func(x float64) color.Color {
		return color.RGBA{
			R: colorSin(x + (math.Pi * 0 / 3)),
			G: colorSin(x + (math.Pi * 2 / 3)),
			B: colorSin(x + (math.Pi * 4 / 3)),
			A: 255,
		}
	}
	return ColorizerFunc(func(res mandelbrot.Result) color.Color {
		if res.Bounded {
			return color.Black
		}
		return colorRainbow(float64(res.Iter) / 4)
	})
}
