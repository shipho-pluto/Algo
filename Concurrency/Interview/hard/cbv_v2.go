package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	BatchV2 struct {
		mu         sync.Mutex
		wg         *sync.WaitGroup
		chRes      chan []any
		cntBatch   int
		batchSlice []any
	}
)

func NewBatchV2(cntBatch int) *BatchV2 {
	return &BatchV2{
		chRes:      make(chan []any),
		wg:         new(sync.WaitGroup),
		mu:         sync.Mutex{},
		cntBatch:   cntBatch,
		batchSlice: make([]any, 0, cntBatch),
	}
}

func (b *BatchV2) doBatching(ctx context.Context, c chan any) chan []any {
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case val, ok := <-c:
			if !ok {
				break LOOP
			}
			b.wg.Add(1)
			go func() {
				defer b.wg.Done()

				b.mu.Lock()
				b.batchSlice = append(b.batchSlice, val)
				if len(b.batchSlice) == b.cntBatch {
					select {
					case <-ctx.Done():
						return
					case b.chRes <- b.batchSlice:
					}
					b.batchSlice = make([]any, 0, b.cntBatch)
				}
				b.mu.Unlock()
			}()
		}
	}

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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	chIn := make(chan any)

	go func() {
		defer close(chIn)
		for i := range 50 {
			chIn <- i + 1
		}
	}()

	b := NewBatchV2(3)
	chOut := b.doBatching(ctx, chIn)

	for val := range chOut {
		fmt.Print(val, " ")
	}
}
