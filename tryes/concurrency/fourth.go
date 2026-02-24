package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func Get2(ctx context.Context, address string, key string) Response4 {
	time.Sleep(1 * time.Second)
	return Response4{Request4(address), nil}
}

func main() {
	addresses := []string{
		"127.0.0.1", "127.0.0.2",
		"127.0.0.3", "127.0.0.4",
		"127.0.0.5", "127.0.0.6",
		"127.0.0.7", "127.0.0.8",
		"127.0.0.9", "127.0.0.10",
	}
	retry := 4
	waitingTime := 5
	nums := 6
	fmt.Println(witchAnsFasterLoop(waitingTime, addresses, retry, nums))
}

type (
	WorkerPool4 struct {
		chanOut     chan Response4
		chanErr     chan error
		start       int
		loop        int
		waitingTime int
	}
	Request4 string

	Response4 struct {
		value Request4
		err   error
	}
)

func (p *WorkerPool4) Start(ctx context.Context, addresses []string) string {
	for _, addr := range addresses {
		go func() {
			res := Get2(ctx, addr, "")
			if res.err != nil {
				p.chanErr <- res.err
			} else {
				p.chanOut <- res
			}
		}()
	}

	select {
	case res := <-p.chanOut:
		if p.start == p.loop {
			return string(res.value)
		}
		return string(res.value) + " with left retries: " + strconv.Itoa(p.loop)
	case <-ctx.Done():
		if p.loop == 0 {
			return ctx.Err().Error()
		}
		newctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.waitingTime*2)*time.Second)
		defer cancel()
		newp := New4(p.loop-1, p.waitingTime*2)
		return newp.Start(newctx, addresses)
	case res := <-p.chanErr:
		return res.Error()
	}
}

func New4(retry, waitingTime int) *WorkerPool4 {
	return &WorkerPool4{
		chanOut:     make(chan Response4),
		loop:        retry,
		start:       retry,
		waitingTime: waitingTime,
	}
}

func witchAnsFasterLoop(waitingTime int, addresses []string, retry, nums int) map[string]int {
	result := make(map[string]int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitingTime)*time.Second)
	defer cancel()
	wp := New4(retry, waitingTime)
	for range nums {
		result[wp.Start(ctx, addresses)]++
	}
	return result
}
