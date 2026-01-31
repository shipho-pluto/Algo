package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randLongCall() int64 {
	rnd := rand.Int63n(2000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

var defaultTimer = 1 * time.Second

func wrapperCall(ctx context.Context) (int64, error) {
	start := time.Now()
	ch := make(chan int64)

	go func() {
		defer close(ch)
		ch <- randLongCall()
	}()

	select {
	case <-ctx.Done():
		fmt.Println(time.Since(start))
		return 0, ctx.Err()
	case val := <-ch:
		fmt.Println(time.Since(start))
		return val, nil
	}
}

func main() {
	fmt.Println("Hello World")
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimer)
	defer cancel()

	res, err := wrapperCall(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
