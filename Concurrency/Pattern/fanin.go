package main

import (
	"context"
	"sync"
)

func fanIn(ctx context.Context, chans ...<-chan int) chan<- int {
	result := make(chan int)
	wg := sync.WaitGroup{}

	for _, job := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case v, ok := <-job:
					if !ok {
						return
					}
					select {
					case result <- v:
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
		close(result)
	}()

	return result
}
