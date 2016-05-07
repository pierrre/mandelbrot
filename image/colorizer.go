package image

import (
	"image/color"

	"github.com/pierrre/mandelbrot"
)

type Colorizer func(complex128, mandelbrot.Result) color.Color

func ColorColorizer(col color.Color) Colorizer {
	return func(c complex128, res mandelbrot.Result) color.Color {
		return col
	}
}

func ColorsIterColorizer(cols []color.Color, shift int) Colorizer {
	return func(c complex128, res mandelbrot.Result) color.Color {
		return cols[(res.Iter+shift)%len(cols)]
	}
}

func BoundColorizer(bounded, unbounded Colorizer) Colorizer {
	return func(c complex128, res mandelbrot.Result) color.Color {
		if res.Bounded {
			return bounded(c, res)
		} else {
			return unbounded(c, res)
		}
	}
}

func BWColorizer(invert bool) Colorizer {
	bounded := ColorColorizer(color.White)
	unbounded := ColorColorizer(color.Black)
	if invert {
		bounded, unbounded = unbounded, bounded
	}
	return BoundColorizer(bounded, unbounded)
}
