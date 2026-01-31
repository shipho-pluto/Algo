package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func workerPool[T any](ctx context.Context, in <-chan T, transform func(T) T, workerNum int) <-chan T {
	out := make(chan T)
	wg := &sync.WaitGroup{}

	wg.Add(workerNum)
	for range workerNum {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- transform(v):
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	urls := make([]string, 1_000_000)

	chIn := make(chan string)
	go func() {
		defer close(chIn)
		for i := range urls {
			chIn <- strconv.Itoa(i)
		}
	}()

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	var cntGoroutines = runtime.GOMAXPROCS(0)
	mapCode := make(map[int]int)

	for range cntGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range chIn {
				code := sendRequest2(url)

				mu.Lock()
				mapCode[code]++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(mapCode)
}

func sendRequest2(url string) (code int) {
	val, _ := strconv.Atoi(url)
	return val
}
