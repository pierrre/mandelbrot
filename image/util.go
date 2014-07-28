package image

import (
	"image"
	"math"
)

func MaxIter(scale float64) uint {
	return uint(math.Log(scale) * 4.5)
}

func ImageScale(im image.Image) float64 {
	size := im.Bounds().Size()
	return math.Min(float64(size.X), float64(size.Y)) / 4
}
