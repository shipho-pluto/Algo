package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

type MyError struct {
	v string
}

func (e *MyError) Error() string {
	return e.v
}

func main() {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		// выполнение работы 1
		return nil
	})

	g.Go(func() error {
		// выполнение работы 2
		someError := MyError{v: "some error"}
		fmt.Println(someError.Error())
		return nil
	})

	if err := g.Wait(); err != nil {
		// обработка ошибки
	}
}
