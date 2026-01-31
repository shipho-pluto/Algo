package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	Batch struct {
		chRes      chan []any
		batchSlice []any
		batchSize  int
		wg         sync.WaitGroup
	}
)

func New(batchSize int) *Batch {
	return &Batch{
		chRes:      make(chan []any, batchSize),
		batchSlice: make([]any, 0, batchSize),
		batchSize:  batchSize,
		wg:         sync.WaitGroup{},
	}
}

func (b *Batch) doBatching(ctx context.Context, c chan any) chan []any {
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		for i := 0; ; {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-c:
				if !ok {
					return
				}
				b.batchSlice = append(b.batchSlice, val)
				i++
			}

			if len(b.batchSlice) == b.batchSize {
				i = 0
				select {
				case <-ctx.Done():
					return
				case b.chRes <- b.batchSlice:
				}
				b.batchSlice = make([]any, 0, b.batchSize)
			}
		}
	}()

	go func() {
		b.wg.Wait()
		if len(b.batchSlice) > 0 {
			b.chRes <- b.batchSlice
		}
		close(b.chRes)
	}()

	return b.chRes
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chIn := make(chan any, 10)

	go func() {
		defer close(chIn)
		for i := range 5 {
			chIn <- i + 1
		}
	}()

	b := New(2)
	chOut := b.doBatching(ctx, chIn)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range chOut {
			fmt.Print(val, " ")
		}
	}()

	wg.Wait()
}
