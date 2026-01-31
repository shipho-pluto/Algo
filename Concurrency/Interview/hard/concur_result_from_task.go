package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type (
	Task struct {
		id   int
		desc string
	}
	Result struct {
		task Task
		err  error
	}
	WorkerPool struct {
		chanIn     chan Task
		workerNums int
		wg         sync.WaitGroup
		chanOut    chan Result
	}
)

var (
	ErrorTimeLimitToRead  = errors.New("time limit to read")
	ErrorTimeLimitToWrite = errors.New("time limit to write")
)

func NewWorkerPool(workerNums int) *WorkerPool {
	return &WorkerPool{
		chanIn:     make(chan Task),
		chanOut:    make(chan Result),
		workerNums: workerNums,
		wg:         sync.WaitGroup{},
	}
}

func (pool *WorkerPool) Submit(ctx context.Context, tasks []Task, transform func(Task) Result) {
	go func() {
		defer close(pool.chanIn)
		for _, task := range tasks {
			select {
			case pool.chanIn <- task:
			case <-ctx.Done():
				pool.chanIn <- Task{id: task.id, desc: "ErrorTimeLimitToWrite"}
			}
		}
	}()

	for range pool.workerNums {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()
			for {
				select {
				case task, ok := <-pool.chanIn:
					if !ok {
						return
					}
					if task.desc == "ErrorTimeLimitToWrite" {
						pool.chanOut <- Result{task: Task{id: task.id}, err: ErrorTimeLimitToWrite}
					} else {
						pool.chanOut <- pool.wrapper(ctx, transform, task)
					}
				}
			}
		}()
	}

	go func() {
		pool.wg.Wait()
		close(pool.chanOut)
	}()
}

func (pool *WorkerPool) wrapper(ctx context.Context, transform func(Task) Result, task Task) Result {
	result := make(chan Result)
	go func() {
		defer close(result)
		result <- transform(task)
	}()

	select {
	case <-ctx.Done():
		return Result{task: Task{id: task.id}, err: ErrorTimeLimitToRead}
	case value := <-result:
		return value
	}
}

func main() {
	tasks := []Task{
		{1, "1"}, {2, "2"},
		{3, "3"}, {4, "4"},
		{5, "5"}, {6, "6"},
		{7, "7"}, {8, "8"},
		{9, "9"}, {10, "10"},
	}
	workerNum := 5
	funcGetResult := func(task Task) Result {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		return Result{task: task, err: nil}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool := NewWorkerPool(workerNum)
	go pool.Submit(ctx, tasks, funcGetResult)

	for v := range pool.chanOut {
		if v.err != nil {
			fmt.Println("ID", v.task.id, "\t", "Error:", v.err)
		} else {
			fmt.Println("ID", v.task.id, "\t", v.task.desc)
		}
	}
}
