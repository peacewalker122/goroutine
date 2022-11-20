package syncpackage

import (
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWaitGroup(t *testing.T) {
	var new Coroutines
	// file, err := os.Create("sync.txt")
	// require.NoError(t, err)
	val := []string{"https://open.spotify.com/", "https://www.mediawiki.org/wiki/API:Main_page"}
	for _, v := range val {
		new = Newg(&sync.WaitGroup{}, v, 1, 0)
	}
	done := make(chan interface{})
	new.PrimeCounter(done)
}

func TestMutex(t *testing.T) {
	strings := "piho"
	log.Println(len(strings))
	require.Equal(t, 42, len(strings))
}

func BenchmarkWaitGroup(b *testing.B) {
	var news Coroutines
	var count <-chan int
	var done chan interface{}
	var Primecounts int

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		//circle, _ = twod.GetShape(twod.Circlestring, float64(i))\\
		wg := new(sync.WaitGroup)
		news = Newg(wg, "", Primecounts, i)
		count, Primecounts = news.PrimeCounter(done)
		b.StopTimer()
		for v := range count {
			Primecounts = v
			//MutexLock(v, &sync.Mutex{})
			fmt.Println("result: ", v)
		}
	}

}

func BenchmarkFetchConcurent(b *testing.B) {
	wg := new(sync.WaitGroup)
	var news Coroutines
	val := []string{"https://open.spotify.com/", "https://www.mediawiki.org/wiki/API:Main_page"}
	done := make(chan interface{})
	for _, v := range val {
		news = Newg(wg, v, 0, 1)
		resp := news.CoroutinesFetch(done)
		fmt.Printf("response: %v \n", <-resp)
	}

}
