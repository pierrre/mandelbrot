package image

type Projection interface {
	Project(x, y int) complex128
}

type ProjectionFunc func(x, y int) complex128

func (pf ProjectionFunc) Project(x, y int) complex128 {
	return pf(x, y)
}
