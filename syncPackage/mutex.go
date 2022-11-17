package syncpackage

import (
	"fmt"
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

func (n *mutex) Helloworld() error {
	n.Lock()
	fmt.Println("HelloWorld")
	defer n.Unlock()

	return nil
}

