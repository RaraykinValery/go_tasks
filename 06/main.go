package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)

	// Горутина, которая читает из канала, пока канал открыт
	go func() {
		fmt.Println("Starting goroutine")

		// Когда горутина пытается прочитать из закрытого канала - цикл прерывается
		for n := range ch {
			fmt.Printf("Reading %v\n", n)
		}

		fmt.Println("Closing goroutine")
		wg.Done()
	}()

	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(time.Second / 2)
	}

	close(ch)

	wg.Wait()
}
