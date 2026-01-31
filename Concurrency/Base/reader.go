package main

import (
	"fmt"
)

func READER(ch <-chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v, "worker1")
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println(v, "worker2")
		}
	}()
}

func main() {
	ch := make(chan int)

	go func() {
		for i := range 20 {
			ch <- i
		}
		close(ch)
	}()

	READER(ch)
}
