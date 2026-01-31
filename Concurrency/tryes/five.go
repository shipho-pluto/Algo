package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
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
	ErrorWait struct {
		chanOut     chan error
		chanIn      chan Input
		chanSign    chan error
		inputData   []Input
		wg          *sync.WaitGroup
		result      error
		maxParallel int
	}

	Input struct {
		ctx context.Context
		foo func(ctx context.Context) error
	}
)

func NewErrorWait(maxParallel int) *ErrorWait {
	return &ErrorWait{
		chanOut:     make(chan error),
		chanSign:    make(chan error),
		chanIn:      make(chan Input),
		inputData:   make([]Input, 0),
		wg:          &sync.WaitGroup{},
		maxParallel: maxParallel,
	}
}

func (e *ErrorWait) run(ctx context.Context, f func(ctx context.Context) error) {
	e.inputData = append(e.inputData, Input{ctx, f})
}

func (e *ErrorWait) wait() error {
	myContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		defer close(e.chanIn)
		for _, input := range e.inputData {
			select {
			case <-myContext.Done():
				return
			case e.chanIn <- input:
			}
		}
	}()

	for range e.maxParallel {
		e.wg.Add(1)
		go func() {
			defer e.wg.Done()
			for {
				select {
				case <-myContext.Done():
					return
				case input, ok := <-e.chanIn:
					if !ok {
						return
					}
					e.wg.Add(1)
					go func() {
						defer e.wg.Done()
						e.chanSign <- input.foo(input.ctx)
					}()

					select {
					case <-myContext.Done():
						return
					case res := <-e.chanSign:
						select {
						case <-myContext.Done():
							return
						case e.chanOut <- res:
						}
					}
				}
			}
		}()
	}

	go func() {
		e.wg.Wait()
		close(e.chanOut)
		close(e.chanSign)
	}()

	for err := range e.chanOut {
		e.result = errors.Join(e.result, err)
	}

	return e.result
}

func main() {
	g := NewErrorWait(2)
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
	if !errors.Is(err, ExpErr1) || !errors.Is(err, ExpErr2) {
		panic(err.Error())
	} else {
		fmt.Println("ok")
	}
}
