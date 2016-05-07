package image

import (
	"image"
	"math/cmplx"
)

type Transformation func(complex128) complex128

func identityTransformation(c complex128) complex128 {
	return c
}

func BaseTransformation(im image.Image, rotate float64, scale float64, translate complex128) Transformation {
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

func ImageTransformation(im image.Image) Transformation {
	center := complex(float64(im.Bounds().Dx())/2, float64(im.Bounds().Dy())/2)
	return func(c complex128) complex128 {
		c -= center
		c = complex(real(c), -imag(c))
		return c
	}
}

func RotateTransformation(rotate float64) Transformation {
	if rotate == 0 {
		return identityTransformation
	}
	rot := cmplx.Rect(1, rotate)
	return func(c complex128) complex128 {
		return c * rot
	}
}

func ScaleTransformation(scale float64) Transformation {
	if scale == 1 {
		return identityTransformation
	}
	invScale := 1 / scale
	return func(c complex128) complex128 {
		return complex(real(c)*invScale, imag(c)*invScale)
	}
}

func TranslateTransformation(translate complex128) Transformation {
	if translate == 0 {
		return identityTransformation
	}
	return func(c complex128) complex128 {
		return c + translate
	}
}
