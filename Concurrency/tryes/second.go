package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := range 10 {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := range 10 {
			ch2 <- i + 10
		}
	}()

	go func() {
		defer close(ch3)
		for i := range 10 {
			ch3 <- i + 20
		}
	}()

	for item := range mergeChan(ch1, ch2, ch3) {
		fmt.Println(item)
	}
}

func mergeChan(chans ...chan int) <-chan int {
	chanIn := make(chan chan int)
	chanOut := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanIn)
		for _, ch := range chans {
			chanIn <- ch
		}
	}()

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ch := range chanIn {
				for val := range ch {
					chanOut <- val
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
