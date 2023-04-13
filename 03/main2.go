package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var sum int64 = 0
	init_arr := []int64{2, 4, 6, 8, 10}

	for _, v := range init_arr {
		wg.Add(1)

		go func(n int64) {
			atomic.AddInt64(&sum, n*n)
			wg.Done()
		}(v)
	}

	wg.Wait()

	fmt.Printf("Sum of squares is %d\n", sum)
}
