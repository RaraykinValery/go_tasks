package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println("Starting goroutine")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closing goroutine")
				wg.Done()
				return
			default:
				fmt.Println("Goroutine doing work...")
				time.Sleep(time.Second / 2)
			}
		}
	}()

	time.Sleep(time.Second)

	cancel()

	wg.Wait()
}
