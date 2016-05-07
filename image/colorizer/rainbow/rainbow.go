package rainbow

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func Colorizer(colorCount int, shift int) mandelbrot_image.Colorizer {
	cols := make([]color.Color, colorCount)
	for i := 0; i < colorCount; i++ {
		cols[i] = colorful.Hsv(float64(i)/float64(colorCount)*360, 1, 1)
	}
	return mandelbrot_image.ColorsIterColorizer(cols, shift)
}
