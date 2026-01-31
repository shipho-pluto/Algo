package main

import (
	"fmt"
	"sync"
)

func tee(in <-chan int, limit int) []<-chan int {
	out := make([]chan int, limit)

	for i := 0; i < limit; i++ {
		out[i] = make(chan int)
	}

	go func() {
		for v := range in {
			for i := range limit {
				out[i] <- v
			}
		}

		for v := range out {
			close(out[v])
		}
	}()

	res := make([]<-chan int, limit)
	for i := 0; i < limit; i++ {
		res[i] = out[i]
	}
	return res
}

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	chs := tee(ch, 2)

	go func() {
		defer wg.Done()
		for val := range chs[0] {
			fmt.Println(val)
		}
	}()

	go func() {
		defer wg.Done()
		for val := range chs[1] {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
