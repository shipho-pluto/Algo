package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	ResponseAddress struct {
		wg       sync.WaitGroup
		chanOut  chan Output
		chanSign chan Output
		chanErr  chan error
		loop     int
	}

	Output struct {
		value string
		err   error
	}
)

func NewRA() *ResponseAddress {
	return &ResponseAddress{
		wg:       sync.WaitGroup{},
		chanOut:  make(chan Output),
		chanSign: make(chan Output),
		chanErr:  make(chan error),
	}
}

func (a *ResponseAddress) Start(ctx context.Context, addresses []string) Output {
	for _, address := range addresses {
		a.wg.Add(1)
		go func() {
			defer a.wg.Done()

			a.wg.Add(1)
			go func() {
				defer a.wg.Done()
				a.chanOut <- Get2(ctx, address, "")
			}()

			select {
			case <-ctx.Done():
				return
			case out := <-a.chanSign:
				if out.err != nil {
					a.chanErr <- out.err
				} else {
					a.chanOut <- out
				}
				a.chanSign = nil
			}
		}()
	}

	go func() {
		a.wg.Wait()
		close(a.chanOut)
		close(a.chanErr)
		if a.chanSign != nil {
			close(a.chanSign)
		}
	}()

	select {
	case <-ctx.Done():
		if a.loop < 5 {
			cg := NewRA()
			cg.loop = a.loop + 1
			newCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return cg.Start(newCtx, addresses)
		}
		return Output{err: ctx.Err()}
	case out := <-a.chanOut:
		return out
	case err := <-a.chanErr:
		return Output{err: err}
	}
}

func main() {
	addresses := []string{
		"127.0.0.1", "127.0.0.2",
		"127.0.0.3", "127.0.0.4",
		"127.0.0.5", "127.0.0.6",
		"127.0.0.7", "127.0.0.8",
		"127.0.0.9", "127.0.0.10",
	}
	result := make(map[string]int)
	for range 20 {
		cg := NewRA()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		result[cg.Start(ctx, addresses).value]++
		cancel()
	}
	fmt.Println(result)
}

func Get2(ctx context.Context, address string, key string) Output {
	time.Sleep(1 * time.Second)
	return Output{address, nil}
}
