package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type (
	ConcurGet struct {
		chanAddress chan string
		chanErr     chan error
		chanOut     chan OutputData
		chanSign    chan OutputData
		maxParallel int
		wg          *sync.WaitGroup
	}

	OutputData struct {
		value string
		err   error
	}
)

func NewCG(maxParallel int) *ConcurGet {
	return &ConcurGet{
		chanAddress: make(chan string),
		chanErr:     make(chan error),
		chanOut:     make(chan OutputData),
		chanSign:    make(chan OutputData),
		maxParallel: maxParallel,
		wg:          &sync.WaitGroup{},
	}
}

func (c *ConcurGet) Start(ctx context.Context, addresses []string) OutputData {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		defer close(c.chanAddress)
		for _, add := range addresses {
			select {
			case <-ctx.Done():
				return
			case c.chanAddress <- add:
			}
		}
	}()

	for range c.maxParallel {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-c.chanAddress:
					if !ok {
						return
					}

					c.wg.Add(1)
					go func() {
						defer c.wg.Done()
						result, err := Get(ctx, data, "")
						if err != nil {
							c.chanSign <- OutputData{err: err}
						} else {
							c.chanSign <- OutputData{value: result}
						}
					}()

					select {
					case <-ctx.Done():
						return
					case outputData, ok := <-c.chanSign:
						if !ok {
							return
						}
						if outputData.err != nil {
							c.chanErr <- outputData.err
						} else {
							c.chanOut <- outputData
						}
						c.chanSign = nil
					}
				}
			}
		}()
	}

	go func() {
		c.wg.Wait()
		close(c.chanOut)
		close(c.chanErr)
	}()

	select {
	case <-ctx.Done():
		return OutputData{err: ErrorNotFound}
	case out := <-c.chanOut:
		return out
	case err := <-c.chanErr:
		return OutputData{err: err}
	}
}

var ErrorNotFound = errors.New("not found")

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
		cg := NewCG(len(addresses))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		result[cg.Start(ctx, addresses).value]++
		cancel()
	}
	fmt.Println(result)
}

func Get(ctx context.Context, address string, key string) (string, error) {
	time.Sleep(5 * time.Second)
	return address, nil
}
