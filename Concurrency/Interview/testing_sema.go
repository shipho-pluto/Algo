package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	{
		sem := make(chan struct{}, 3)
		wg := &sync.WaitGroup{}
		for i := 1; i <= 10; i++ {
			sem <- struct{}{}
			wg.Add(1)
			go Print1(i, sem, wg)
		}
		wg.Wait()
	}
	fmt.Println("---------------------")
	{
		in := make(chan int)
		go func() {
			defer close(in)
			for i := 1; i <= 10; i++ {
				in <- i
			}
		}()
		wg := &sync.WaitGroup{}
		for range 3 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for val := range in {
					Print2(val)
				}
			}()
		}
		wg.Wait()
	}
}

func Print1(i int, sem chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		<-sem
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}

func Print2(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
