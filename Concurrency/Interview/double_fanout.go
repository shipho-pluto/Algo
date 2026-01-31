package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan any)

	go func() {
		defer close(in)
		for i := range 10 {
			in <- i
		}
	}()

	res1, res2 := tee(in)
	go func() {
		for v := range res1 {
			fmt.Println(v)
		}
	}()
	for v := range res2 {
		fmt.Println(v)
	}
}

func tee(in chan any) (chan any, chan any) {
	wg := sync.WaitGroup{}
	out1, out2 := make(chan any), make(chan any)

	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for {
			select {
			case v, ok := <-in:
				if !ok {
					break LOOP
				}
				select {
				case out1 <- v:
					select {
					case out2 <- v:
					default:
					}
				case out2 <- v:
					select {
					case out1 <- v:
					default:
					}
				default:
				}
			default:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out1)
		close(out2)
	}()

	return out1, out2
}
