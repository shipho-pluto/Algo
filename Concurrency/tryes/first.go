package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	WorkerPool struct {
		maxParallelism int
		mu             sync.Mutex
		wg             sync.WaitGroup
		chanIn         chan Request
	}

	Request string
)

func New(maxParallelism int) *WorkerPool {
	return &WorkerPool{
		maxParallelism: maxParallelism,
		chanIn:         make(chan Request),
		mu:             sync.Mutex{},
		wg:             sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start(ctx context.Context, urls []string, codes map[int]int) {
	wp.wg.Add(1)
	go func() {
		defer close(wp.chanIn)
		defer wp.wg.Done()
	LoopIn:
		for _, url := range urls {
			select {
			case <-ctx.Done():
				break LoopIn
			case wp.chanIn <- Request(url):
			}
		}
	}()

	for range wp.maxParallelism {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
		LoopOut:
			for {
				select {
				case <-ctx.Done():
					break LoopOut
				case req, ok := <-wp.chanIn:
					if !ok {
						break LoopOut
					}
					chanSign := make(chan struct{})
					wp.wg.Add(1)
					go func() {
						defer wp.wg.Done()
						defer close(chanSign)
						wp.mu.Lock()
						codes[len(req)]++
						wp.mu.Unlock()
					}()

					select {
					case <-ctx.Done():
						break LoopOut
					case <-chanSign:
					}
				}
			}
		}()
	}

	wp.wg.Wait()
}

func main() {
	urls := make([]string, 1_000_000_000)
	codes := make(map[int]int)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wp := New(1_000)
	wp.Start(ctx, urls, codes)
	fmt.Println(codes)
}
