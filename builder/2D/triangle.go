package builder

import "math"

type triangle[T float64] struct {
	// Length For Every Line In The Triangle including the base.
	Length T
}

func newTriangle[T float64](Length T) twoDimentional[T] {
	// in this block we declaring a contract to the interface
	// implement the "TwoDimentional" interface the "Area" and "Circumfence"

	return &triangle[T]{
		Length: Length,
	}
}

func (n *triangle[T]) Area() T {
	return T((math.Pow(float64(n.Length), 2)) / 2)
}
func (n *triangle[T]) Circumfence() T {
	return T(n.Length * 3)
}
