package main

import (
	"fmt"
	"sync"
)

func WRITER() <-chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 11
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {

	ch := WRITER()

	for v := range ch {
		fmt.Println(v)
	}

}
