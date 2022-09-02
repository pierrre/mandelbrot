package image

import (
	"image"
	"math"
)

// MaxIter returns the maximum iteration required for the given scale.
func MaxIter(scale float64) int {
	return int(math.Log(scale) * 10)
}

// Scale returns the required scale for the given image size.
func Scale(size image.Point) float64 {
	return math.Min(float64(size.X), float64(size.Y)) / 4
}
