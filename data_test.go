package main

import (
	"context"
	"fmt"
	"log"
	file "testfile/TEST"
	"testing"
	"time"
)

type times []<-chan*time.Time

const(
	totalfile = 3000
	contentlenght = 5000
)

func BenchmarkCode(b *testing.B) {
	b.Run("add test", func(b *testing.B) {
		ctx := context.Background()
		c := make(chan *time.Time,2)
		//c1 := make(chan *time.Time, 1)
		p := file.Newperson("s")
		go p.Add(ctx, c)
		go p.Add(ctx, c)
		go p.Add(ctx, c)
		//go p.add(ctx, c1)

		c <- file.RandTime()
		//c1 <- randTime()
		defer close(c)
		//defer close(c1)
		//defer b.Name()
	})
}

func BenchmarkMerge(b *testing.B) {
	b.Run("mergetest", func(b *testing.B) {
		c := file.Filltimep(1)
		c1 := file.Filltimep(1)
		c2 := file.Filltimep(1)
		c3 := file.Filltimep(1)
		c4 := file.Filltimep(1)

		s := file.Newperson("s")
		log.Printf("length is %v", len(c))
		go s.Add(context.Background(), c)

		merge := file.MergeTimeP(context.Background(), c, c1, c2, c3, c4)
		for i := range merge {
			log.Println(i.UTC())
		}
		go s.Add(context.Background(), merge)
	})
}

func BenchmarkConsolidate(b *testing.B) {
	c := file.Filltimep(1)
	c1 := file.Filltimep(1)
	c2 := file.Filltimep(1)
	c3 := file.Filltimep(1)
	c4 := file.Filltimep(1)
	
	p := file.Newperson("S")

	r := times{c,c1,c2,c3,c4}

	s := file.Consolidate(r...)

	for i := range s{
		fmt.Println(i.UTC())
	}
	p.Add(context.Background(),s)
}