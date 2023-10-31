package image

import (
	"image/color"

	"github.com/pierrre/mandelbrot"
)

// Colorizer is a function that returns a color for a point and a result.
type Colorizer func(complex128, mandelbrot.Result) color.Color

// ColorColorizer returns a [Colorizer] that always returns the same color.
func ColorColorizer(col color.Color) Colorizer {
	return func(c complex128, res mandelbrot.Result) color.Color {
		return col
	}
}

// ColorsIterColorizer returns a [Colorizer] that uses a list of colors.
func ColorsIterColorizer(cols []color.Color, shift int) Colorizer {
	return func(c complex128, res mandelbrot.Result) color.Color {
		return cols[(res.Iter+shift)%len(cols)]
	}
}

// BoundColorizer returns a [Colorizer] that uses a bounded color for bounded points and an unbounded color for unbounded points.
func BoundColorizer(bounded, unbounded Colorizer) Colorizer {
	return func(c complex128, res mandelbrot.Result) color.Color {
		if res.Bounded {
			return bounded(c, res)
		}
		return unbounded(c, res)
	}
}

// BWColorizer returns a [Colorizer] that uses a black and white color scheme.
func BWColorizer(invert bool) Colorizer {
	bounded := ColorColorizer(color.White)
	unbounded := ColorColorizer(color.Black)
	if invert {
		bounded, unbounded = unbounded, bounded
	}
	return BoundColorizer(bounded, unbounded)
}
