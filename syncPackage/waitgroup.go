package syncpackage

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Coroutines interface {
	PrimeCounter(done <-chan interface{}) (<-chan int, int)
	CoroutinesFetch(done <-chan interface{}) <-chan *http.Response
}

type wg struct {
	sync  *sync.WaitGroup
	value int
	count int
	url   string
}

func Newg(sync *sync.WaitGroup, url string, count, value int) Coroutines {
	return &wg{
		sync:  sync,
		value: value,
		count: count,
		url:   url,
	}
}

func (n *wg) PrimeCounter(done <-chan interface{}) (<-chan int, int) {
	channel := make(chan int, 1)

	go func() {
		defer close(channel)

		select {
		case channel <- n.PrimeCount():
		case <-time.After(5 * time.Second):
			return
		case <-done:
			return
		}

	}()
	return channel, n.count
}

func (n *wg) CoroutinesFetch(done <-chan interface{}) <-chan *http.Response {
	response := make(chan *http.Response)
	go func() {
		defer close(response)
		resp, err := http.Get(n.url)
		if err != nil {
			fmt.Println(err.Error())
		}
		select {
		case <-done:
			return
		case <-time.After(5 * time.Second):
			return
		case response <- resp:
		}

	}()
	return response
}

func MutexLock[T any](arg T, mutex sync.Locker) {
	mutex.Lock()
	log.Printf("result is: %v", arg)
	defer mutex.Unlock()
}

func (n *wg) PrimeCount() int {

	if n.value == 2 {
		n.count++
	} else if n.value%2 != 0 {
		n.count++
	}

	return n.count
}
