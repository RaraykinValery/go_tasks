package main

import "fmt"

func main() {
	uniq_map := make(map[string]bool)
	finalSet := make([]string, 0)

	strings_slice := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, v := range strings_slice {
		if _, ok := uniq_map[v]; !ok {
			uniq_map[v] = true
		}
	}

	for k := range uniq_map {
		finalSet = append(finalSet, k)
	}

	fmt.Println(finalSet)
}
