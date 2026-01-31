package main

import "fmt"

func transform(in <-chan int, trans func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- trans(v)
		}
	}()

	return out
}

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 5 {
			ch <- i
		}
	}()

	mul := func(v int) int {
		return v * v
	}

	out := transform(ch, mul)
	for v := range out {
		fmt.Println(v)
	}
}
