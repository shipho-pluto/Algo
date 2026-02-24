package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go func() {
		defer close(chan1)
		for i := range 10 {
			chan1 <- i
		}
	}()
	go func() {
		defer close(chan2)
		for i := 10; i < 20; i++ {
			chan2 <- i
		}
	}()
	go func() {
		defer close(chan3)
		for i := 20; i < 30; i++ {
			chan3 <- i
		}
	}()

	for v := range fanIn2(chan1, chan2, chan3) {
		fmt.Println(v)
	}
}

func fanIn2(chans ...chan int) <-chan int {
	res := make(chan int)
	wg := sync.WaitGroup{}
	for _, chanI := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-chanI:
					if !ok {
						return
					}
					res <- val
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
