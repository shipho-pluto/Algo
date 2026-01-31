package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	ConcurBucket struct {
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		chanRes     chan []int
		slice       []int
		cntInBucket int
	}
)

func NewConcurBucket(cntInBkt int) *ConcurBucket {
	return &ConcurBucket{
		wg:          &sync.WaitGroup{},
		mu:          &sync.Mutex{},
		chanRes:     make(chan []int),
		cntInBucket: cntInBkt,
		slice:       make([]int, 0, cntInBkt),
	}
}

func (b *ConcurBucket) Start(ctx context.Context, chanIn <-chan int) chan []int {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case element, ok := <-chanIn:
			if !ok {
				break Loop
			}
			b.wg.Add(1)
			go func() {
				defer b.wg.Done()

				b.mu.Lock()
				b.slice = append(b.slice, element)
				if len(b.slice) == b.cntInBucket {
					select {
					case b.chanRes <- b.slice:
					case <-ctx.Done():
						return
					}
					b.slice = make([]int, 0, b.cntInBucket)
				}
				b.mu.Unlock()
			}()
		}
	}

	go func() {
		b.wg.Wait()
		if len(b.slice) > 0 {
			b.chanRes <- b.slice
		}
		close(b.chanRes)
	}()

	return b.chanRes
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cb := NewConcurBucket(1)

	chIn := make(chan int)
	go func() {
		defer close(chIn)
		for i := range 50 {
			chIn <- i + 1
		}
	}()

	chanOut := cb.Start(ctx, chIn)
	var result [][]int
	for val := range chanOut {
		result = append(result, val)
	}
	fmt.Println(result)
}
