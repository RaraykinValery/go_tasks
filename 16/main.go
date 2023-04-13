package main

import "fmt"

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		var less []int
		var greater []int

		for _, v := range array[1:] {
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
