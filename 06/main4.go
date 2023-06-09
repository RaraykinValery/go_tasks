package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := false

	var wg sync.WaitGroup

	wg.Add(1)

	// Горутина будет работать до тех пор, пока глобальный флаг done имеет значение flase
	go func() {
		fmt.Println("Starting goroutine")

		for {
			if done {
				fmt.Println("Closing goroutine")
				wg.Done()
				return
			}

			fmt.Println("Goroutine doing work...")
			time.Sleep(time.Second / 2)
		}
	}()

	time.Sleep(time.Second)

	done = true

	wg.Wait()
}
