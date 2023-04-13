package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, in <-chan int) {
	for n := range in {
		fmt.Printf("Worker #%d prints: %d\n", id, n)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	permanent_flow_chan := make(chan int)

	workers_num := flag.Int("workers", 1, "number of workers")
	flag.Parse()

	for i := 1; i <= *workers_num; i++ {
		wg.Add(1)
		go func(n int) {
			worker(n, permanent_flow_chan)
			fmt.Printf("Worker #%d stoped\n", n)
			wg.Done()
		}(i)
	}

	i := 0

loop:
	for {
		select {
		case <-quit:
			fmt.Println("Exiting the programm...")
			close(permanent_flow_chan)
			wg.Wait()
			break loop
		default:
			permanent_flow_chan <- i
			i++
		}
	}
}
