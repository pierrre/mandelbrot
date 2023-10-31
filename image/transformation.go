package image

import (
	"image"
	"math/cmplx"
)

// Transformation is a transformation function.
type Transformation func(complex128) complex128

func identityTransformation(c complex128) complex128 {
	return c
}

// BaseTransformation returns a [Transformation] function for the given parameters.
func BaseTransformation(im image.Image, rotate, scale float64, translate complex128) Transformation {
	it := ImageTransformation(im)
	rt := RotateTransformation(rotate)
	st := ScaleTransformation(scale)
	tt := TranslateTransformation(translate)
	return func(c complex128) complex128 {
		c = it(c)
		c = rt(c)
		c = st(c)
		c = tt(c)
		return c
	}
}

// ImageTransformation returns a [Transformation] function for the given image.
func ImageTransformation(im image.Image) Transformation {
	center := complex(float64(im.Bounds().Dx())/2, float64(im.Bounds().Dy())/2)
	return func(c complex128) complex128 {
		c -= center
		c = complex(real(c), -imag(c))
		return c
	}
}

// RotateTransformation returns a [Transformation] function for the given rotation.
func RotateTransformation(rotate float64) Transformation {
	if rotate == 0 {
		return identityTransformation
	}
	rot := cmplx.Rect(1, rotate)
	return func(c complex128) complex128 {
		return c * rot
	}
}

// ScaleTransformation returns a [Transformation] function for the given scale.
func ScaleTransformation(scale float64) Transformation {
	if scale == 1 {
		return identityTransformation
	}
	invScale := 1 / scale
	return func(c complex128) complex128 {
		return complex(real(c)*invScale, imag(c)*invScale)
	}
}

// TranslateTransformation returns a [Transformation] function for the given translation.
func TranslateTransformation(translate complex128) Transformation {
	if translate == 0 {
		return identityTransformation
	}
	return func(c complex128) complex128 {
		return c + translate
	}
}
