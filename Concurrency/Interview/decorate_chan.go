package main

import (
	"fmt"
	"sync"
)

type respAlg struct {
	id  int
	val int
}

func decorate(id int, f func(int) int, jobs <-chan int, results chan<- respAlg) {
	wg := sync.WaitGroup{}

	wg.Add(id)
	for i := range id {
		go func() {
			defer wg.Done()
			for v := range jobs {
				results <- respAlg{id: i, val: f(v)}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}()

	res := make(chan respAlg)

	mul := func(x int) int {
		return x * x
	}

	go decorate(5, mul, ch, res)

	for v := range res {
		fmt.Println(v)
	}
}
