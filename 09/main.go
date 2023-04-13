package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	singles := make(chan int)
	doubles := make(chan int)

	init_arr := []int{2, 4, 6, 7, 8}

	wg.Add(2)

	// Горутина, удваивающая исходные числа
	// Закрывает канал doubles, чтобы завершить горутину, которая из него читает
	go func() {
		for v := range singles {
			doubles <- v * 2
		}

		close(doubles)

		wg.Done()
	}()

	// Пишет удвоенные значения в stdout
	// После закрытия канала doubles горутина завершается
	go func() {
		for v := range doubles {
			fmt.Println(v)
		}

		wg.Done()
	}()

	// Пишем числа в singles и закрываем его, чтобы завершить горутину, из него читающую
	for _, v := range init_arr {
		singles <- v
	}

	close(singles)

	wg.Wait()
}
