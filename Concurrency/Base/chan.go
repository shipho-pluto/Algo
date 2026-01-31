package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	{
		ch := make(chan int)

		for i := range 5 {
			go func() {
				ch <- i + 1
			}()
		}

		for range 1 {
			val := <-ch

			fmt.Println(val)
		}
	}

	{
		ch := make(chan int)
		initSeconds := time.Now()

		for range 100 {
			go func() {
				ch <- randomWait()
			}()
		}

		totalSeconds := 0
		for range 100 {
			totalSeconds += <-ch
		}

		mainSeconds := time.Since(initSeconds).Seconds()
		fmt.Println(totalSeconds, mainSeconds)
	}

	{
		wg := &sync.WaitGroup{}
		mt := &sync.Mutex{}
		totalSeconds := 0
		wg.Add(100)
		for range 100 {
			go func() {
				defer wg.Done()
				seconds := randomWait()
				mt.Lock()
				totalSeconds += seconds
				mt.Unlock()

			}()
		}
		wg.Wait()

		fmt.Println(totalSeconds)
	}
}

var maxSeconds = 5

func randomWait() int {
	workSeconds := rand.Intn(maxSeconds + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}
