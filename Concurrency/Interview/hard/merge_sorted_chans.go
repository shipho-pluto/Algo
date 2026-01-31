package main

import (
	"fmt"
	"sync"
)

func mergeSortedChans(ch1, ch2 <-chan int) <-chan int {
	var slice1, slice2 []int
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			slice1 = append(slice1, v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch2 {
			slice2 = append(slice2, v)
		}
	}()

	wg.Wait()
	out := make(chan int)

	go func() {
		defer close(out)
		for l, r := 0, 0; l+r < len(slice1)+len(slice2); {
			if l < len(slice1) && slice1[l] < slice2[r] {
				out <- slice1[l]
				l++
			} else if r < len(slice2) {
				out <- slice2[r]
				r++
			}
		}
	}()

	return out
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		defer close(ch1)
		for i := range 5 {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := range 5 {
			ch2 <- i + 5
		}
	}()

	for num := range mergeSortedChans(ch1, ch2) {
		fmt.Println(num)
	}
}
