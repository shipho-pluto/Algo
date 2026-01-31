package main

import (
	"fmt"
	"sync"
)

func main() {
	var money int
	var donation int

	mutex := &sync.RWMutex{}

	go func() {
		for {
			mutex.RLock()
			m := money
			d := donation
			mutex.RUnlock()

			if m != d {
				fmt.Println("money", m, d)
				break
			}
		}
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()

			mutex.Lock()
			money++
			donation++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(money)
}
