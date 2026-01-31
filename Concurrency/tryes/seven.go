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

	for val := range sortedMergeChan(ch1, ch2) {
		fmt.Println(val)
	}
}

func sortedMergeChan(ch1 <-chan int, ch2 <-chan int) <-chan int {
	wg := sync.WaitGroup{}
	var sl1, sl2 []int

	wg.Add(1)
	go func() {
		defer wg.Done()
		for el := range ch1 {
			sl1 = append(sl1, el)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for el := range ch2 {
			sl2 = append(sl2, el)
		}
	}()

	wg.Wait()
	chanOut := make(chan int)

	go func() {
		defer close(chanOut)
		for l, r := 0, 0; l < len(sl1) || r < len(sl2); {
			if l < len(sl1) && r < len(sl2) {
				if sl1[l] < sl2[r] {
					chanOut <- sl1[l]
					l++
				} else {
					chanOut <- sl2[r]
					r++
				}
			} else if l < len(sl1) {
				chanOut <- sl1[l]
				l++
			} else {
				chanOut <- sl2[r]
				r++
			}
		}
	}()

	return chanOut
}
