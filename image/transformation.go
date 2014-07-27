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

var IdentityTransformation = TransformationFunc(func(c complex128) complex128 {
	return c
})

func BaseTransformation(im image.Image, scale float64, translate complex128) Transformation {
	var trans Transformation
	trans = ImageTransformation(im)
	trans = ScaleTransformation(scale, trans)
	trans = TranslateTransformation(translate, trans)
	return trans
}

func ImageTransformation(im image.Image) Transformation {
	center := complex(float64(im.Bounds().Dx())/2, -float64(im.Bounds().Dy())/2)
	return TransformationFunc(func(c complex128) complex128 {
		c = complex(real(c), -imag(c))
		c -= center
		return c
	})
}

func ScaleTransformation(scale float64, trans Transformation) Transformation {
	return TransformationFunc(func(c complex128) complex128 {
		c = trans.Transform(c)
		return complex(real(c)/scale, imag(c)/scale)
	})
}

func TranslateTransformation(translate complex128, trans Transformation) Transformation {
	return TransformationFunc(func(c complex128) complex128 {
		return trans.Transform(c) + translate
	})
}
