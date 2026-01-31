package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch1 <- i
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch2 <- i + 10
		}
		close(ch2)
	}()

	go func() { wg.Wait() }()

	select {
	case v, ok := <-ch1:
		if ok {
			fmt.Println(v, "case1")
		}
	case v, ok := <-ch2:
		if ok {
			fmt.Println(v, "case2")
		}
	}

	{
		cnt, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan int)
		go func() {
			for i := range 10 {
				select {
				case ch <- i:
				case <-cnt.Done():
				}
			}
			close(ch)
		}()

		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			case <-cnt.Done():
				return
			}

		}
	}
}
