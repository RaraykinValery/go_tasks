package main

import "fmt"

func main() {
	intersection_map := make(map[int]bool)
	intersection := make([]int, 0)

	set1 := []int{1, 5, 4, 3, 2}
	set2 := []int{4, 3, 5, 7, 6, 8}

	for _, v := range set1 {
		intersection_map[v] = false
	}

	for _, v := range set2 {
		if _, ok := intersection_map[v]; ok {
			intersection_map[v] = true
		}
	}

	for k, v := range intersection_map {
		if v {
			intersection = append(intersection, k)
		}
	}

	fmt.Println(intersection)
}
