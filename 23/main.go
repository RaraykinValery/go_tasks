package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 3

	fmt.Println(s)

	s = append(s[:i], s[i+1:]...)

	fmt.Println(s)
}
