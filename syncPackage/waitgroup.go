package syncpackage

import (
	"sync"
)

type Coroutines[T any] interface {
	Helloworld() T
}

type wg[T any] struct {
	sync  *sync.WaitGroup
	value []T
}

func Newg[T any](sync *sync.WaitGroup, value ...T) Coroutines[T] {
	return &wg[T]{
		sync:  sync,
		value: value,
	}
}

func (n *wg[T]) Helloworld() T {
	channel := make(chan T, len(n.value))
	for _, i := range n.value {
		n.sync.Add(1)
		go func() {
			defer n.sync.Done()
			channel <- i
		}()
	}
	n.sync.Wait()
	return T(<-channel)
}
