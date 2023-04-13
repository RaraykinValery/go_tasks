package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	words := strings.Fields(s)
	reversed_fields := make([]string, len(words))

	for i, w := range words {
		reversed_fields[len(words)-1-i] = w
	}

	reversed := strings.Join(reversed_fields, " ")

	fmt.Println(reversed)
}
