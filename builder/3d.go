package builder

type ThreeDimentional[T float64] interface {
	Volume() T
}

type Tube[T float64] struct {
	TwoDimentional TwoDimentional[T]
	Height         T
}

func NewTube[T float64](Height, Radius T) ThreeDimentional[T] {
	circle := NewCircle(Radius)
	return &Tube[T]{
		TwoDimentional: circle,
		Height:         Height,
	}
}

func (n *Tube[T]) Volume() T {
	return T(n.TwoDimentional.Area() * n.Height)
}
