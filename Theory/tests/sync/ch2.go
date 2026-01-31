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

type outData struct {
	val string
	err error
}

type inData struct {
	val string
	i   int
}

var ErrorCxtTimeOut = errors.New("time out by context")

func processData(cxt context.Context, val string) <-chan outData {
	ch := make(chan int)
	out := make(chan outData)

	go func() {
		defer close(ch)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}()

	go func() {
		select {
		case <-cxt.Done():
			out <- outData{"", ErrorCxtTimeOut}
		case <-ch:
			out <- outData{"url " + val + " OK", nil}
		}
	}()

	return out
}

func helper(cxt context.Context, urls []string, limit int) []http.Response {
	output := make([]http.Response, len(urls))

	in := make(chan inData)
	out := make(chan inData)

	go func() {
		defer close(in)
		for i := range len(urls) {
			select {
			case <-cxt.Done():
				return
			case in <- inData{urls[i], i}:
			}
		}
	}()

	mainTime := time.Now()
	precessByLimit(cxt, in, out, limit)

	var results []inData
	for v := range out {
		results = append(results, v)
		fmt.Println("Value =", v.val)
	}

	fmt.Println(time.Since(mainTime))

	sort.Slice(results, func(i, j int) bool {
		return results[i].i < results[j].i
	})
	for i := 0; i < len(output); i++ {
		if i < len(results) {
			output[i] = http.Response{
				Status: results[i].val,
			}
		} else {
			output[i] = http.Response{
				Status: ErrorCxtTimeOut.Error(),
			}
		}
	}

	return output
}

func precessByLimit(cxt context.Context, in <-chan inData, out chan<- inData, limit int) {
	wg := &sync.WaitGroup{}
	for range limit {
		wg.Add(1)
		go worker(cxt, in, out, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func worker(cxt context.Context, in <-chan inData, out chan<- inData, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-cxt.Done():
			return
		case val, ok := <-in:
			if !ok {
				return
			}
			select {
			case <-cxt.Done():
				return
			case res := <-processData(cxt, val.val):
				if errors.Is(res.err, ErrorCxtTimeOut) {
					return
				}
				select {
				case <-cxt.Done():
					return
				case out <- inData{res.val, val.i}:
				}
			}
		}
	}
}

func main() {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := []string{"1", "2", "5", "4", "3"}
	limit := 5

	output := helper(cxt, input, limit)

	for i := range output {
		fmt.Print(output[i].Status, " ")
	}
}
