package image

type Transformation interface {
	Transform(complex128) complex128
}

type TransformationFunc func(complex128) complex128

func (f TransformationFunc) Transform(c complex128) complex128 {
	return f(c)
}
