package main

import (
	"context"
	"fmt"
	"sync"
)

type (
	WorkerChan struct {
		outCh chan interface{}
		wg    sync.WaitGroup
	}
)

func NewWorkerChan() *WorkerChan {
	return &WorkerChan{
		outCh: make(chan interface{}),
		wg:    sync.WaitGroup{},
	}
}

func (wc *WorkerChan) Bridge(ctx context.Context, genChan <-chan <-chan interface{}) <-chan interface{} {
	go func() {
		defer func() {
			wc.wg.Wait()
			close(wc.outCh)
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case chIn, ok := <-genChan:
				wc.wg.Add(1)
				if !ok {
					wc.wg.Done()
					return
				}
				go func() {
					defer wc.wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						case val, ok := <-chIn:
							if !ok {
								return
							}
							select {
							case <-ctx.Done():
								return
							case wc.outCh <- val:
							}
						}
					}
				}()
			}
		}
	}()

	return wc.outCh
}

func main() {
	out := make(chan (<-chan interface{}))
	go func() {
		defer close(out)
		for j := range 10 {
			stream := make(chan interface{}, 10)
			for i := range 10 {
				stream <- i + 10*j
			}
			close(stream)
			out <- stream
		}
	}()

	var res []interface{}
	w := NewWorkerChan()
	for v := range w.Bridge(context.Background(), out) {
		res = append(res, v)
	}

	fmt.Println(res)
}
