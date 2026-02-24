package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ExpErr1 = errors.New("got error 1")
	ExpErr2 = errors.New("got error 2")
	ExpErr3 = errors.New("got error 3")
)

type (
	Waiter interface {
		wait() error
		run(ctx context.Context, f func(ctx context.Context) error)
	}

	MyWaiter struct {
		saveFunc    []Data
		chanIn      chan Data
		mt          sync.Mutex
		wgIn        sync.WaitGroup
		wgOut       sync.WaitGroup
		maxParallel int
	}

	Data struct {
		v   func(ctx context.Context) error
		ctx context.Context
	}
)

func NewWaiter(parallel int) *MyWaiter {
	return &MyWaiter{
		chanIn:      make(chan Data),
		wgIn:        sync.WaitGroup{},
		wgOut:       sync.WaitGroup{},
		mt:          sync.Mutex{},
		maxParallel: parallel,
	}
}

func (w *MyWaiter) run(ctx context.Context, f func(ctx context.Context) error) {
	w.saveFunc = append(w.saveFunc, Data{f, ctx})
}

func (w *MyWaiter) wait() error {

	for _, f := range w.saveFunc {
		w.wgIn.Add(1)
		go func() {
			defer w.wgIn.Done()
			select {
			case w.chanIn <- f:
			}
		}()
	}

	go func() {
		w.wgIn.Wait()
		defer close(w.chanIn)
	}()

	rErr := make(chan error)
	for range w.maxParallel {
		w.wgOut.Add(1)
		go func() {
			defer w.wgOut.Done()
			for {
				select {
				case val, ok := <-w.chanIn:
					if !ok {
						return
					}

					w.wgOut.Add(1)
					var err error
					signChan := make(chan struct{})
					go func() {
						defer w.wgOut.Done()
						defer close(signChan)
						err = val.v(val.ctx)
					}()

					select {
					case <-signChan:
						select {
						case rErr <- err:
						case <-val.ctx.Done():
							return
						}
					case <-val.ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		w.wgOut.Wait()
		close(rErr)
	}()

	var res error
	for err := range rErr {
		res = errors.Join(res, err)
	}
	return res
}

func main() {
	g := NewWaiter(2)
	ctx := context.Background()

	g.run(ctx, func(ctx context.Context) error {
		return ExpErr1
	})

	g.run(ctx, func(ctx context.Context) error {
		return ExpErr2
	})

	g.run(ctx, func(ctx context.Context) error {
		return nil
	})

	err := g.wait()
	if !(errors.Is(err, ExpErr2) && errors.Is(err, ExpErr1)) {
		panic("something went wrong")
	} else {
		fmt.Println("ok")
	}

}
