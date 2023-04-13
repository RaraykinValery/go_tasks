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

	// Горутина, которая читает из канала, пока он не будет закрыт
	go func() {
		// Когда горутина пытается прочитать из закрытого канала - цикл прерывается
		// После выхода из цикла горутина останавливается
		fmt.Println("Starting goroutine")

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
