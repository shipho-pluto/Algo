package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.reddit.com",
		"https://www.twitter.com",
		"https://www.youtube.com",
		"https://www.youtube", // Некорректный URL (вызовет ошибку)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	errChan := make(chan error, 1)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			err := fetchUrl(url)
			if err != nil {
				select {
				case errChan <- err:
					cancel()
				default:
				}
				return
			}

			select {
			case <-ctx.Done():
				return
			default:
			}
			fmt.Printf("Fetched %s\n", url)
		}(url)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()
	if err := <-errChan; err != nil {
		fmt.Printf("Stopped due to error: %v\n", err)
	} else {
		fmt.Println("All requests completed successfully!")
	}
}

func fetchUrl(url string) error {
	_, err := http.Get(url)
	return err
}
