package main

import (
	"fmt"
	"sync"
)

type Sum struct {
	Value int
	mu    sync.Mutex
}

func (s *Sum) incr(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Value += n
}

var wg sync.WaitGroup

func main() {
	sum := Sum{}
	init_arr := []int{2, 4, 6, 8, 10}

	for _, v := range init_arr {
		wg.Add(1)
		go func(n int) {
			sum.incr(n * n)
			wg.Done()
		}(v)
	}

	wg.Wait()

	fmt.Printf("Sum of squares is %d\n", sum.Value)
}
