package image

import (
	"image"
	"math"
)

type Transformation interface {
	Transform(complex128) complex128
}

type TransformationFunc func(complex128) complex128

func (f TransformationFunc) Transform(c complex128) complex128 {
	return f(c)
}

var identityTransformation = TransformationFunc(func(c complex128) complex128 {
	return c
})

type ListTransformation []Transformation

func (lt ListTransformation) Transform(c complex128) complex128 {
	for _, trans := range lt {
		c = trans.Transform(c)
	}
	return c
}

func Transformations(transformations ...Transformation) Transformation {
	return ListTransformation(transformations)
}

func BaseTransformation(im image.Image, rotate float64, scale float64, translate complex128) Transformation {
	return Transformations(
		ImageTransformation(im),
		RotateTransformation(rotate),
		ScaleTransformation(scale),
		TranslateTransformation(translate),
	)
}

func ImageTransformation(im image.Image) Transformation {
	center := complex(float64(im.Bounds().Dx())/2, -float64(im.Bounds().Dy())/2)
	return TransformationFunc(func(c complex128) complex128 {
		c = complex(real(c), -imag(c))
		c -= center
		return c
	})
}

func RotateTransformation(rotate float64) Transformation {
	if rotate == 0 {
		return identityTransformation
	}
	return TransformationFunc(func(c complex128) complex128 {
		return complex(
			real(c)*math.Cos(rotate)-imag(c)*math.Sin(rotate),
			real(c)*math.Sin(rotate)+imag(c)*math.Cos(rotate),
		)
	})
}

func ScaleTransformation(scale float64) Transformation {
	if scale == 1 {
		return identityTransformation
	}
	return TransformationFunc(func(c complex128) complex128 {
		return complex(real(c)/scale, imag(c)/scale)
	})
}

func TranslateTransformation(translate complex128) Transformation {
	if translate == 0 {
		return identityTransformation
	}
	return TransformationFunc(func(c complex128) complex128 {
		return c + translate
	})
}
