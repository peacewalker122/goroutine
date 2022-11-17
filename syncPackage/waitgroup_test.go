package syncpackage

import (
	"log"
	"os"
	"strconv"
	"sync"
	twod "testfile/builder/2D"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWaitGroup(t *testing.T) {
	file, err := os.Create("sync.txt")
	require.NoError(t, err)
	val := []string{"HELLO WORLD", "HALO DUNIA", "HAI DUNIA"}
	new := Newg(&sync.WaitGroup{}, file, val...)

	new.Helloworld()
}

func TestMutex(t *testing.T) {
	mutex := NewMutex(&sync.WaitGroup{}, &sync.Mutex{})
	err := mutex.Helloworld()
	require.NoError(t, err)
}

func BenchmarkWaitGroup(b *testing.B) {
	var circle twod.Shape
	var val []string

	file, err := os.Create("sync.txt")
	if err != nil {
		log.Panic(err.Error())
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		circle, err = twod.GetShape(twod.Circlestring, float64(i))
		MutexLock(circle.Area(), &sync.Mutex{})

		if err != nil {
			log.Panic(err.Error())
		}
		res := strconv.Itoa(int(circle.Area()))
		val = append(val, res)
	}

	new := Newg(&sync.WaitGroup{}, file, val...)
	err = new.Helloworld()
	b.StopTimer()
	if err != nil {
		log.Panic(err.Error())
	}
}
