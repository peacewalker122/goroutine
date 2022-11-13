package file

import "sync"

type Concurent[T any] struct{
	param *sync.WaitGroup
	value T
}

func newConcurent[T any](param *sync.WaitGroup,value T) *Concurent[T]{
	return &Concurent[T]{
		param: param,
		value:     value,
	}
}

func(s *Concurent[T])