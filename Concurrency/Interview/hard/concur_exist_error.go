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
)

type (
	Waiter interface {
		wait() error
		run(ctx context.Context, f func(ctx context.Context) error)
	}
	WaitGroupStruct struct {
		wg          sync.WaitGroup
		inputSlice  []InputData
		inputChan   chan InputData
		errChan     chan error
		result      error
		maxParallel int
	}

	InputData struct {
		ctx context.Context
		foo func(ctx context.Context) error
	}
)

func NewGroupWait(maxParallel int) Waiter {
	return &WaitGroupStruct{
		wg:          sync.WaitGroup{},
		inputSlice:  []InputData{},
		inputChan:   make(chan InputData),
		errChan:     make(chan error),
		maxParallel: maxParallel,
	}
}

func (g *WaitGroupStruct) run(ctx context.Context, f func(ctx context.Context) error) {
	g.inputSlice = append(g.inputSlice, InputData{ctx: ctx, foo: f})
}

func (g *WaitGroupStruct) wait() error {
	go func() {
		defer close(g.inputChan)
		for _, input := range g.inputSlice {
			g.inputChan <- input
		}
	}()

	for range g.maxParallel {
		g.wg.Add(1)
		go func() {
			defer g.wg.Done()
			for {
				select {
				case input, ok := <-g.inputChan:
					if !ok {
						return
					}
					ctx, foo := input.ctx, input.foo
					signChan := make(chan error)
					go func() {
						defer close(signChan)
						signChan <- foo(ctx)
					}()

					select {
					case <-ctx.Done():
						return
					case err := <-signChan:
						g.errChan <- err
					}
				}
			}
		}()
	}

	go func() {
		for err := range g.errChan {
			g.result = errors.Join(g.result, err)
		}
	}()

	g.wg.Wait()
	close(g.errChan)

	return g.result
}

func main() {
	g := NewGroupWait(2)
	ctx := context.Background()

	g.run(ctx, func(ctx context.Context) error {
		return ExpErr2
	})

	g.run(ctx, func(ctx context.Context) error {
		return ExpErr1
	})

	g.run(ctx, func(ctx context.Context) error {
		return nil
	})

	err := g.wait()

	if !errors.Is(err, ExpErr1) || !errors.Is(err, ExpErr2) {
		panic(err.Error())
	} else {
		fmt.Println("ok")
	}
}
