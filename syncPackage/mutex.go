package syncpackage

import (
	"fmt"
	"net/http"
	"sync"
)

type mutex struct {
	*sync.WaitGroup
	*sync.Mutex
}

func NewMutex(wg *sync.WaitGroup, mutexs *sync.Mutex) Coroutines {
	return &mutex{
		WaitGroup: wg,
		Mutex:     mutexs,
	}
}

func (n *mutex) PrimeCounter(done <-chan interface{}) (<-chan int, int) {
	n.Lock()
	fmt.Println("HelloWorld")
	defer n.Unlock()

	return nil, 1
}

func (n *mutex) CoroutinesFetch(done <-chan interface{}) <-chan *http.Response {
	panic("unin")
}
