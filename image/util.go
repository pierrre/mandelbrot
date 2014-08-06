package image

import (
	"image"
	"math"
)

func MaxIter(scale float64) int {
	return int(math.Log(scale) * 10)
}

func ImageScale(size image.Point) float64 {
	return math.Min(float64(size.X), float64(size.Y)) / 4
}
