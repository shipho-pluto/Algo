package main

import (
	"fmt"
	"sync"
)

func parse(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func send(in <-chan int, limit int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for range limit {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range in {
				out <- v
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
	ch := make(chan int)
	limit := 2

	go func() {
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}()

	for value := range send(parse(ch), limit) {
		fmt.Println(value)
	}
}
