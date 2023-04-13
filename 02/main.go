package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	init_arr := []int{2, 4, 6, 8, 10}

	for _, v := range init_arr {
		wg.Add(1)

		go func(n int) {
			fmt.Printf("%d^2 = %d\n", n, n*n)
			wg.Done()
		}(v)
	}

	wg.Wait()
}
