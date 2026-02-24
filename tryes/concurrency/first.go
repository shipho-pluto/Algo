package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	urls := make([]string, 1_000_000_000)
	urls[0] = "www.youtube.com"
	urls[1] = "https://google.com"
	urls[2] = "www.twitch.io"
	urls[3] = "http://localhost:8080"

	codes := make(map[int]int)
	parallel := 1_000

	fmt.Println(full(urls, codes, parallel))
}

type (
	WorkerPull1 struct {
		chanIn      chan Request1
		maxParallel int
		wg          sync.WaitGroup
		mt          sync.Mutex
	}

	Request1 string
)

func New1(parallel int) *WorkerPull1 {
	return &WorkerPull1{
		chanIn:      make(chan Request1),
		maxParallel: parallel,
		wg:          sync.WaitGroup{},
		mt:          sync.Mutex{},
	}
}

func (wp *WorkerPull1) Start(ctx context.Context, urls []string, codes map[int]int) {
	wp.wg.Add(1)
	go func() {
		defer close(wp.chanIn)
		defer wp.wg.Done()

		for _, url := range urls {
			select {
			case wp.chanIn <- Request1(url):
			case <-ctx.Done():
				return
			}
		}
	}()

	for range wp.maxParallel {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case req, ok := <-wp.chanIn:
					if !ok {
						return
					}
					signChan := make(chan struct{})
					wp.wg.Add(1)
					go func() {
						defer close(signChan)
						defer wp.wg.Done()

						wp.mt.Lock()
						codes[len(req)]++
						wp.mt.Unlock()
					}()

					select {
					case <-ctx.Done():
						return
					case <-signChan:
					}
				}
			}
		}()
	}

	wp.wg.Wait()
}

func full(urls []string, codes map[int]int, parallel int) any {
	wp := New1(parallel)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wp.Start(ctx, urls, codes)

	return codes
}
