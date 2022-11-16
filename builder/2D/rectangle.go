package builder

import "math"

type rectangle[T float64] struct {
	Length T
}

func newRectangle[T float64](Length T) twoDimentional[T] {
	// in this block we declaring a contract to the interface
	// implement the "TwoDimentional" interface the "Area" and "Circumfence"
	return &rectangle[T]{
		Length: Length,
	}
}

func (n *rectangle[T]) Area() T {
	length := float64(n.Length)

	return T(math.Pow(length, 2))
}
func (n *rectangle[T]) Circumfence() T {
	length := float64(n.Length)

	return T(2 * length)
}
