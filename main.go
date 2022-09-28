package main

import (
	"context"
	"log"
	file "testfile/TEST"
)

func main() {
	ctx := context.Background()
	//done:= make(chan struct{})
	r := file.Newperson("s")
	// cs := make(chan *time.Time,1)
	// cs1 := make(chan *time.Time,1)

	c := file.Filltimep(1)
	c1 := file.Filltimep(1)
	c3 := file.Filltimep(1)
	c4 := file.Filltimep(1)
	t1 := file.Filltime()
	t2 := file.Filltime()

	out := file.MergeTimeP(context.Background(), c, c1, c3, c4)
	for a := range out {
		log.Printf("result: %v", a.UTC())
	}

	go r.Add(ctx, out)

	outs := file.MergeTime(context.Background(), t1, t2)
	for i := range outs {
		log.Println(i.UTC())
	}
	// for{
	// go func() {
	// 	c <- randTime()
	// 	c1 <- randTime()
	// 	defer close(c)
	// 	defer close(c1)
	// }()
	// 	select{
	// 	case <-done:
	// 		return
	// 	case <-time.After(200 * time.Millisecond):
	// 		log.Println("TimeOut")
	// 		panic("timeout")
	// 	}
	// }
}
