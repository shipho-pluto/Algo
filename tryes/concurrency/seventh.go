package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

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

	for i := range sortedChan(ch1, ch2) {
		fmt.Println(i)
	}
}

func sortedChan(ch1 chan int, ch2 chan int) <-chan int {
	res := make(chan int)
	wg := sync.WaitGroup{}

	var arr2, arr1 []int
	for i := range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if i == 0 {
				for v := range ch1 {
					arr1 = append(arr1, v)
				}
			} else {
				for v := range ch2 {
					arr2 = append(arr2, v)
				}
			}
		}()
	}

	wg.Wait()

	go func() {
		defer close(res)
		for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
			if l < len(arr1) && (r < len(arr2) && arr1[l] < arr2[r] || r == len(arr2)) {
				res <- arr1[l]
				l++
			} else {
				res <- arr2[r]
				r++
			}
		}
	}()

	return res
}
