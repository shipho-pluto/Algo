package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	WareHouse struct {
		items     map[string]int
		mu        sync.RWMutex
		wg        sync.WaitGroup
		cntFailed atomic.Int32
	}
	Order struct {
		ID       int
		Item     string
		Quantity int
	}
)

func NewWareHouse() *WareHouse {
	return &WareHouse{
		items: map[string]int{
			"phone": 4, "laptop": 3,
			"tablet": 3, "tv": 2,
			"computer": 1, "ps5": 3,
			"mouse": 6, "headphone": 7,
			"keyboard": 4, "monitor": 3,
		},
		mu: sync.RWMutex{},
		wg: sync.WaitGroup{},
	}
}

func (w *WareHouse) ReserveItem(item string, qnt int) bool {
	w.mu.RLock()
	if val, ok := w.items[item]; val-qnt < 0 && ok {
		w.mu.RUnlock()
		return false
	}
	w.mu.RUnlock()

	w.mu.Lock()
	w.items[item] -= qnt
	w.mu.Unlock()
	return true
}

func ProcessOrders(ctx context.Context, inChan <-chan Order, maxParallel int) {
	w := NewWareHouse()

	for range maxParallel {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case order, ok := <-inChan:
					if !ok {
						return
					}

					signal := make(chan struct{})

					w.wg.Add(1)
					go func() {
						defer w.wg.Done()
						defer close(signal)
						val := w.ReserveItem(order.Item, order.Quantity)
						if val {
							fmt.Println("Successful reserve order:", order)
						} else {
							fmt.Println("Failed reserve order:", order)
							w.cntFailed.Add(1)
						}
					}()

					select {
					case <-ctx.Done():
						return
					case <-signal:
					}
				}
			}
		}()
	}

	w.wg.Wait()
	fmt.Println("Count of failed orders:", w.cntFailed.Load())
}

func main() {
	orders := make(chan Order)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		defer close(orders)
		for i := range 10 {
			item := []string{"phone", "laptop", "tablet", "tv", "computer",
				"ps5", "mouse", "headphone", "keyboard", "monitor"}[rand.Intn(10)]
			qty := rand.Intn(3) + 1
			orders <- Order{Item: item, Quantity: qty, ID: i}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	ProcessOrders(ctx, orders, 3)
	fmt.Println("All orders are complete")
}
