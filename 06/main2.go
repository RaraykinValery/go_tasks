package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)

	// Для остановки горутины можно читать сигнал из канала,
	// предназначенного для предупреждения о конце выполнения программы
	go func() {
		fmt.Println("Starting goroutine")

		for {
			select {
			case <-done:
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

	done <- true

	wg.Wait()
}
