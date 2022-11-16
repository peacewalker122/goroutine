package builder

import (
	"fmt"
	"strings"
)

type twoDimentional[T float64] interface {
	Area() T
	Circumfence() T
}

const (
	Circlestring    = "circle"
	Rectanglestring = "rectangle"
	Trianglestring  = "triangle"
)

// in this block we create a builder to build complex methods for each shape.
func GetShape[T float64](shape string, arg T) (twoDimentional[T], error) {
	// this code below with purpose to getting less error
	res := strings.ToLower(shape)

	if res == Circlestring {
		return newCircle(arg), nil
	}
	if res == Rectanglestring {
		return newRectangle(arg), nil
	}
	if res == Trianglestring {
		return newTriangle(arg), nil
	}
	return nil, fmt.Errorf("unrecognized shape")
}

type diagonal[T float64] struct {
	TwoDimentionals twoDimentional[T]
	diagonal        T
}

func NewDiagonal[T float64](shape twoDimentional[T]) *diagonal[T] {
	// Consume The Shape of the TwoDimentional interface then return into the diagonal struct value
	// Singleton for the
	return &diagonal[T]{
		TwoDimentionals: shape,
	}
}
func (n *diagonal[T]) SetShape(shape twoDimentional[T]) {
	n.TwoDimentionals = shape
}

func (n *diagonal[T]) GetDiagonal() T {
	n.diagonal = n.TwoDimentionals.Area() / n.TwoDimentionals.Circumfence()
	return n.diagonal
}
