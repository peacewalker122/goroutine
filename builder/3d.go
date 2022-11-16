package builder

import twod "testfile/builder/2D"

type ThreeDimensional[T float64] interface {
	Volume() T
}

type Tube[T float64] struct {
	Area        T
	Circumfence T
	Height      T
}

func NewTube[T float64](Height, Radius T) (ThreeDimensional[T], error) {
	circle, err := twod.GetShape("circle", Radius)
	if err != nil {
		return nil, err
	}
	return &Tube[T]{
		Area:        circle.Area(),
		Circumfence: circle.Circumfence(),
		Height:      Height,
	}, nil

}

func (n *Tube[T]) Volume() T {
	return T(n.Area * n.Height)
}
