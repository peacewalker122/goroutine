package builder

import "math"

type circle[T float64] struct {
	Radius T
}

func newCircle[T float64](Radius T) twoDimentional[T] {
	// in this block we declaring a contract to the interface
	// implement the "TwoDimentional" interface the "Area" and "Circumfence"
	return &circle[T]{
		Radius: Radius,
	}
}

func (n *circle[T]) Area() T {
	radius := float64(n.Radius)
	return T(math.Pow(radius, 2)) * math.Pi
}
func (n *circle[T]) Circumfence() T {
	radius := float64(n.Radius)
	return T(2 * (radius * math.Pi))
}
