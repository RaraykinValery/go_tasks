package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	ns := flag.Int("delay", 1, "delay in seconds")
	flag.Parse()

	go func() {
		for n := range ch {
			fmt.Printf("Reading %d\n", n)
		}
	}()

	timeout := time.After(time.Duration(*ns) * time.Second)

	i := 0

	for {
		select {
		case <-timeout:
			fmt.Println("Time is out")
			close(ch)
			return
		default:
			ch <- i
			fmt.Printf("Writting %d\n", i)
			i++
			time.Sleep(time.Second)
		}
	}
}
