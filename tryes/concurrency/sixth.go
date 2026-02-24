package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 50
	buffer := 7
	fmt.Println(generateRandom(n, buffer))
}

func generateRandom(n int, buffer int) [][]int {
	l := n / buffer
	if n%buffer != 0 {
		l++
	}

	res := make([][]int, l)
	chIn := make(chan int)
	wg := sync.WaitGroup{}

	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			chIn <- i
		}()
	}

	for i := range l {
		wg.Add(1)
		go func() {
			defer wg.Done()
			arr := make([]int, 0, buffer)
			for v := range chIn {
				arr = append(arr, v)
				if len(arr) == buffer || i == l-1 && len(arr) == l%buffer {
					res[i] = arr
					return
				}
			}
		}()
	}

	wg.Wait()
	close(chIn)

	return res
}
