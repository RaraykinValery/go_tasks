package main

import (
	"fmt"
	"math/rand"
)

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		randIndex := rand.Intn(len(array))
		pivot := array[randIndex]

		var less []int
		var greater []int

		for i, v := range array {
			if i == randIndex {
				continue
			}

			if v < pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}

		return append(append(quicksort(less), pivot), quicksort(greater)...)
	}
}

func main() {
	numbers := []int{2, 4, 6, 9, 19, 3}

	fmt.Println(numbers)

	numbers = quicksort(numbers)

	fmt.Println(numbers)
}
