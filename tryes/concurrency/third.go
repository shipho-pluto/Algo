package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func Get(ctx context.Context, address Request3, key string) Response3 {
	time.Sleep(250 * time.Millisecond)
	num, err := strconv.Atoi(string(address[len(address)-1]))
	if err != nil {
		return Response3{"", nil}
	} else if num%2 != 0 {
		return Response3{address, ErrorRequest}
	}
	return Response3{address, nil}
}

func main() {
	addrs := []string{
		"127.0.0.1", "127.0.0.2",
		"127.0.0.3", "127.0.0.4",
		"127.0.0.5", "127.0.0.6",
		"127.0.0.7", "127.0.0.8",
		"127.0.0.9", "127.0.0.10",
	}

	res := make(map[string]int)
	for range 20 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		res[witchAnsFaster(ctx, addrs)]++
		cancel()
	}

	fmt.Println(res)
}

type (
	WorkerPool3 struct {
		chanIn  chan Request3
		chanOut chan Response3
		chanErr chan error
	}

	Request3  string
	Response3 struct {
		value Request3
		err   error
	}
)

var ErrorRequest = errors.New("address has odd number")

func (p *WorkerPool3) Start(ctx context.Context, addrs []string) string {
	for _, addr := range addrs {
		go func() {
			select {
			case p.chanIn <- Request3(addr):
			case <-ctx.Done():
				return
			}
		}()
	}

	for range len(addrs) {
		go func() {
			for {
				select {
				case addr, ok := <-p.chanIn:
					if !ok {
						return
					}
					signChan := make(chan struct{})
					go func() {
						defer close(signChan)
						req := Get(ctx, addr, "")
						if req.err != nil {
							p.chanErr <- req.err
						} else {
							p.chanOut <- req
						}
					}()

					select {
					case <-signChan:
						return
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	select {
	case res := <-p.chanOut:
		return string(res.value)
	case err := <-p.chanErr:
		return err.Error()
	case <-ctx.Done():
		return ctx.Err().Error()
	}
}

func New3() *WorkerPool3 {
	return &WorkerPool3{
		chanIn:  make(chan Request3),
		chanOut: make(chan Response3),
		chanErr: make(chan error),
	}
}

func witchAnsFaster(ctx context.Context, addrs []string) string {
	wp := New3()
	return wp.Start(ctx, addrs)
}
