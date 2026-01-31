package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make([]int, 1_000_000)
	transform := func(k int) int {
		return k * k
	}

	semaphore := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for num := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}
			fmt.Println(transform(num))
			<-semaphore
		}()
	}

	wg.Wait()
}
