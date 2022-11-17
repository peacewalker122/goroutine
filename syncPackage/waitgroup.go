package syncpackage

import (
	"log"
	"os"
	"sync"
)

type Coroutines interface {
	Helloworld() error
}

type wg struct {
	sync  *sync.WaitGroup
	file  *os.File
	value []string
}

func Newg(sync *sync.WaitGroup, file *os.File, value ...string) Coroutines {
	return &wg{
		sync:  sync,
		file:  file,
		value: value,
	}
}

func (n *wg) Helloworld() error {
	length := len(n.value)
	channel := make(chan error, length)
	go func() {
		defer close(channel)
		for _, i := range n.value {
			n.sync.Add(1)
			defer n.sync.Done()
			_, err := n.file.WriteString(i)
			channel <- err
		}
	}()
	n.sync.Wait()
	return <-channel
}

func MutexLock[T any](arg T,mutex sync.Locker) {
	mutex.Lock()
	log.Println(arg)
	defer mutex.Unlock()
}
