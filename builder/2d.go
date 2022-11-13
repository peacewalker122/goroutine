package builder

import "math"

type TwoDimentional[T float64] interface {
	Area() T
	Circumfence() T
}

type Circle[T float64] struct {
	Radius T
}

type Rectangle[T float64] struct {
	Length T
}

// in this block we declaring a contract to the interface
// implement the "TwoDimentional" interface the "Area" and "Circumfence"
func NewCircle[T float64](Radius T) TwoDimentional[T] {
	return &Circle[T]{
		Radius: Radius,
	}
}

func NewRectangle[T float64](Length T) TwoDimentional[T] {
	return &Rectangle[T]{
		Length: Length,
	}
}

func (n *Circle[T]) Area() T {
	radius := float64(n.Radius)
	return T(math.Pow(radius, 2)) * math.Pi
}
func (n *Circle[T]) Circumfence() T {
	radius := float64(n.Radius)
	return T(2 * (radius * math.Pi))
}

func (n *Rectangle[T]) Area() T {
	length := float64(n.Length)

	return T(math.Pow(length, 2))
}
func (n *Rectangle[T]) Circumfence() T {
	length := float64(n.Length)

	return T(2 * length)
}
