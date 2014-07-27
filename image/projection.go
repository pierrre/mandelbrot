package image

type Projection interface {
	Project(complex128) complex128
}

type ProjectionFunc func(complex128) complex128

func (f ProjectionFunc) Project(c complex128) complex128 {
	return f(c)
}
