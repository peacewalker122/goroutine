package file

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type any interface{}

type maps map[string]any

type Person struct {
	s  string
	Wg *sync.WaitGroup
}

func Filltimep(size int) <-chan *time.Time {
	c := make(chan *time.Time)
	for {

		go func() {
			c <- RandTime()
			defer close(c)
		}()
		return c

	}
}

func Filltime() <-chan time.Time {
	c := make(chan time.Time)
	for {

		go func() {
			c <- *RandTime()
			defer close(c)
		}()
		select {
		default:
			return c
		case <-time.After(200 * time.Millisecond):
			log.Println("TIMEOOOOOUT")
			return nil
		}

	}
}

func Newperson(s string) *Person {
	return &Person{s: s}
}

func (p *Person) Add(ctx context.Context, c <-chan *time.Time) {
	timeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	var count int
	defer cancel()
	go func() {
		for {
			select {
			case d, ok := <-c:
				if !ok {
					log.Printf("status: %v", ok)
					return
				}
				log.Printf("result%v: %v", p.s, d.UTC())

			case <-timeout.Done():
				count++
				log.Printf("count: %v", count)
				return
			}
		}
	}()
}

func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	select {
	default:
		return out
	case <-time.After(200 * time.Millisecond):
		log.Println("TIMEOOOOOUT")
		return nil
	}
}

func RandTime() *time.Time {
	rand := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randNow := time.Unix(rand, 0)
	return &randNow
}
func randInt() int {
	return rand.Int()
}
func CreateChannels(number, fill int) (chans []<-chan int) {
	chans = make([]<-chan int, number)
	for i := 0; i < number; i++ {
		chans[i] = FillChan(fill)
	}
	return
}
func FillChan(number int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < number; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func MergeTime(ctx context.Context, cs ...<-chan time.Time) <-chan time.Time {
	out := make(chan time.Time, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan time.Time) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func MergeTimeP(ctx context.Context, cs ...<-chan *time.Time) <-chan *time.Time {
	_, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	out := make(chan *time.Time)
	var wg sync.WaitGroup

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan *time.Time) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func FaninTime(c1, cs <-chan *time.Time) <-chan *time.Time {
	out := make(chan *time.Time)
	go func() {
		for {
			select {
			case s := <-cs:
				out <- s
			case s := <-c1:
				out <- s
			case <-time.After(200 * time.Millisecond):
				out <- nil
				return
			}

		}
	}()
	return out
}

func RandomString(length int) string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

    return string(b)
}
