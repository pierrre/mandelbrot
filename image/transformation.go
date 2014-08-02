package image

import (
	"image"
)

type Transformation interface {
	Transform(complex128) complex128
}

type TransformationFunc func(complex128) complex128

func (f TransformationFunc) Transform(c complex128) complex128 {
	return f(c)
}

func BaseTransformation(im image.Image, scale float64, translate complex128) Transformation {
	imageTrans := ImageTransformation(im)
	scaleTrans := ScaleTransformation(scale)
	translateTrans := TranslateTransformation(translate)
	return TransformationFunc(func(c complex128) complex128 {
		c = imageTrans.Transform(c)
		c = scaleTrans.Transform(c)
		c = translateTrans.Transform(c)
		return c
	})
}

func ImageTransformation(im image.Image) Transformation {
	center := complex(float64(im.Bounds().Dx())/2, -float64(im.Bounds().Dy())/2)
	return TransformationFunc(func(c complex128) complex128 {
		c = complex(real(c), -imag(c))
		c -= center
		return c
	})
}

func ScaleTransformation(scale float64) Transformation {
	return TransformationFunc(func(c complex128) complex128 {
		return complex(real(c)/scale, imag(c)/scale)
	})
}

func TranslateTransformation(translate complex128) Transformation {
	return TransformationFunc(func(c complex128) complex128 {
		return c + translate
	})
}
