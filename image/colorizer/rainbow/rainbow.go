// Package rainbow provides a colorizer that uses a rainbow color scheme.
package rainbow

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

// Colorizer returns a [mandelbrot_image.Colorizer] that uses a rainbow color scheme.
func Colorizer(colorCount int, shift int) mandelbrot_image.Colorizer {
	cols := make([]color.Color, colorCount)
	for i := range colorCount {
		cols[i] = colorful.Hsv(float64(i)/float64(colorCount)*360, 1, 1)
	}
	return mandelbrot_image.ColorsIterColorizer(cols, shift)
}
