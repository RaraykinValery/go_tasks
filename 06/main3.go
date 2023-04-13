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

	// Горутина работает до тех пор, пока созданный контекст не будет
	// отменён вызовом cancel(), так как после отмены контекста передаётся сигнал
	// в канал ctx.Done()
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
