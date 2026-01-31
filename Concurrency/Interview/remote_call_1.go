package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type response1 struct {
	id  int
	err error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	chanForResp := make(chan response1)
	go RPCCall(ctx, chanForResp)

	result := <-chanForResp
	if result.err != nil {
		fmt.Println(result.err)
		return
	}
	fmt.Println(result.id)
}

func RPCCall(ctx context.Context, chanForResp chan response1) {
	randInt := rand.Intn(5000)
	ch := make(chan struct{})

	go func() {
		defer close(ch)
		time.Sleep(time.Duration(randInt) * time.Millisecond)
	}()

	select {
	case <-ctx.Done():
		chanForResp <- response1{id: -1, err: ctx.Err()}
		return
	case <-ch:
		chanForResp <- response1{id: randInt}
	}
}
