package main

import (
	"context"
	"sync"
)

func fanOut(ctx context.Context, in <-chan int, numChan int) []chan<- int {
	chans := make([]chan<- int, numChan)
	wg := &sync.WaitGroup{}

	for i := 0; i < numChan; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(chans[i])
			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					select {
					case chans[i] <- v:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
	}()

	return chans
}
