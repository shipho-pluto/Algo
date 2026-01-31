package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Cache struct {
	cache float64
	t     *time.Ticker
}

func NewCache(timer int) *Cache {
	return &Cache{
		t:     time.NewTicker(time.Duration(timer) * time.Second),
		cache: getFloat(),
	}
}

func (c *Cache) updateCourse(ctx context.Context) {
	go func() {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-c.t.C:
				c.cache = getFloat()
			}
		}
	}()
}

func getFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*50_000 + 30_000
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewCache(10)
	c.updateCourse(ctx)

	http.HandleFunc("/getBitcoin", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Bitcoin value is %.2f", c.cache)
		if err != nil {
			log.Print(err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
