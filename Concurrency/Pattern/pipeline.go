package main

import "fmt"

func generate(values ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range values {
			out <- v
		}
	}()

	return out
}

func process(in <-chan int, transform func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- transform(v)
		}
	}()

	return out
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mul := func(v int) int {
		return v * v
	}

	for v := range process(generate(values...), mul) {
		fmt.Println(v)
	}
}
