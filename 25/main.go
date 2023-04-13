package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	startTime := time.Now()
	for time.Since(startTime) < duration {
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start")
	sleep(3 * time.Second)
	fmt.Println("End")
}
