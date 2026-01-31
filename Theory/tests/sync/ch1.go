package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1") // 2
		wg.Done()
	}()

	go func() {
		fmt.Println("2") // 1
	}()
	wg.Wait()
	fmt.Println("3") // 3

	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

LOOP:
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				break LOOP
			}
			fmt.Println("v", v)
		}
	}

}
