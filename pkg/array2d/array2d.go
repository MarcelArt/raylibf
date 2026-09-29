package array2d

type Array2D[T any] struct {
	a    []T
	x, y uint
}

func NewArray2D[T any](x, y uint) *Array2D[T] {
	length := x * y
	return &Array2D[T]{
		a: make([]T, length),
	}
}

func (a Array2D[T]) Idx(x, y uint) uint {
	return x + y*a.x
}

func (a Array2D[T]) Get(x, y uint) T {
	return a.a[a.Idx(x, y)]
}

func (a *Array2D[T]) Set(x, y uint, value T) {
	a.a[a.Idx(x, y)] = value
}
