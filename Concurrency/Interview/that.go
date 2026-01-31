package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"sync"
	"time"
)

var ErrWithContext = errors.New("context cancelled")

type resStruct struct {
	val string
	err error
	id  int
}

func precessData(ctx context.Context, value resStruct) <-chan resStruct {
	ch := make(chan int)
	out := make(chan resStruct)
	go func() {
		defer close(ch)
		time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	}()

	go func() {
		select {
		case <-ctx.Done():
			out <- resStruct{err: ErrWithContext}
		case <-ch:
			out <- resStruct{val: "url " + value.val + " OK", id: value.id, err: nil}
		}
	}()

	return out
}

func starter(ctx context.Context, inputData []string, limit int) []http.Response {
	in := make(chan resStruct)
	out := make(chan resStruct)

	go func() {
		defer close(in)
		for i, val := range inputData {
			select {
			case <-ctx.Done():
				return
			case in <- resStruct{val: val, err: nil, id: i}:
			}
		}
	}()

	start := time.Now()
	precessByLimit(ctx, in, out, limit)

	var outputData []resStruct
	for val := range out {
		fmt.Println(val)
		outputData = append(outputData, val)
	}
	fmt.Println(time.Since(start))

	sort.Slice(outputData, func(i, j int) bool {
		return outputData[i].id < outputData[j].id
	})

	var result []http.Response
	for i := range outputData {
		result = append(result, http.Response{Status: outputData[i].val})

	}
	return result
}

func precessByLimit(ctx context.Context, in <-chan resStruct, out chan<- resStruct, limit int) {
	wg := &sync.WaitGroup{}

	for range limit {
		wg.Add(1)
		go work(ctx, in, out, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func work(ctx context.Context, in <-chan resStruct, out chan<- resStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case value := <-precessData(ctx, v):
				if errors.Is(v.err, ErrWithContext) {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- value:
				}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := []string{"1", "2", "5", "4", "3"}
	limit := 5

	output := starter(ctx, input, limit)
	for val := range output {
		fmt.Println(output[val].Status)
	}
}
