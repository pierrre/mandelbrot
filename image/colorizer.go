package image

import (
	"image/color"

	"github.com/pierrre/mandelbrot"
)

type Colorizer interface {
	Colorize(c complex128, res mandelbrot.Result) color.Color
}

type ColorizerFunc func(c complex128, res mandelbrot.Result) color.Color

func (f ColorizerFunc) Colorize(c complex128, res mandelbrot.Result) color.Color {
	return f(c, res)
}

func ColorColorizer(col color.Color) Colorizer {
	return ColorizerFunc(func(c complex128, res mandelbrot.Result) color.Color {
		return col
	})
}

func ColorsIterColorizer(cols []color.Color, shift int) Colorizer {
	return ColorizerFunc(func(c complex128, res mandelbrot.Result) color.Color {
		return cols[(res.Iter+shift)%len(cols)]
	})
}

func BoundColorizer(bounded, unbounded Colorizer) Colorizer {
	return ColorizerFunc(func(c complex128, res mandelbrot.Result) color.Color {
		if res.Bounded {
			return bounded.Colorize(c, res)
		} else {
			return unbounded.Colorize(c, res)
		}
	})
}

func BWColorizer(invert bool) Colorizer {
	bounded := ColorColorizer(color.White)
	unbounded := ColorColorizer(color.Black)
	if invert {
		bounded, unbounded = unbounded, bounded
	}
	return BoundColorizer(bounded, unbounded)
}
