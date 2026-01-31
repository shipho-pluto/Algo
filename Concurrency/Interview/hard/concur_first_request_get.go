package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func Get(ctx context.Context, address, key string) ResultAddr {
	time.Sleep(2 * time.Second)
	return ResultAddr{value: address, err: nil}
}

var (
	ErrNotFound = errors.New("not found")
	ErrTimeout  = errors.New("timeout")
)

type (
	ResultAddr struct {
		value string
		err   error
	}
	Wrapper struct {
		chanRes chan ResultAddr
		chanErr chan error
		chanSig chan ResultAddr
		wg      sync.WaitGroup
		loop    int
	}
)

func NewWrapper() *Wrapper {
	return &Wrapper{
		chanRes: make(chan ResultAddr),
		chanErr: make(chan error),
		chanSig: make(chan ResultAddr),
		wg:      sync.WaitGroup{},
		loop:    1,
	}
}

func (ww *Wrapper) Get(ctx context.Context, addresses []string, key string) ResultAddr {
	ww.wg.Add(len(addresses))
	for _, address := range addresses {
		go func() {
			defer ww.wg.Done()

			ww.wg.Add(1)
			go func() {
				defer ww.wg.Done()
				ww.chanSig <- Get(ctx, address, key)
			}()

			select {
			case <-ctx.Done():
			case result := <-ww.chanSig:
				if errors.Is(result.err, ErrNotFound) {
					ww.chanErr <- result.err
				} else {
					ww.chanRes <- result
				}
				ww.chanSig = nil
			}
		}()
	}

	go func() {
		ww.wg.Wait()
		if ww.chanSig != nil {
			close(ww.chanSig)
		}
		close(ww.chanRes)
		close(ww.chanErr)
	}()

	select {
	case <-ww.chanErr:
		return ResultAddr{err: ErrNotFound}
	case <-ctx.Done():
		if ww.loop < 5 {
			newCtx, cancel := context.WithTimeout(context.Background(), time.Duration(ww.loop*5)*time.Second)
			defer cancel()

			newWw := NewWrapper()
			newWw.loop += ww.loop
			fmt.Println("Retry", ww.loop)
			return newWw.Get(newCtx, addresses, key)
		}
		return ResultAddr{err: ErrTimeout}
	case result := <-ww.chanRes:
		return result
	}
}

func main() {
	address := []string{
		"127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4", "127.0.0.5",
		"127.0.0.6", "127.0.0.7", "127.0.0.8", "127.0.0.9", "127.0.0.0",
	}
	key := "key"

	cntAddr := make(map[string]int, 10)
	for range 5 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		wrapper := NewWrapper()
		value := wrapper.Get(ctx, address, key)

		if errors.Is(value.err, ErrTimeout) {
			fmt.Println(ErrTimeout.Error())
		} else if errors.Is(value.err, ErrNotFound) {
			fmt.Println(ErrNotFound.Error())
		} else {
			cntAddr[value.value]++
			fmt.Println(value)
		}
		cancel()
	}

	if len(cntAddr) != 0 {
		fmt.Println(cntAddr)
	}
}
