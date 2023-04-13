package main

import "fmt"

func binarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	myList := []int{1, 3, 5, 7, 9}
	item := 5

	result := binarySearch(myList, item)

	if result == -1 {
		fmt.Println("Элемент не найден в массиве")
	} else {
		fmt.Println(result)
	}
}
