package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var changed_int int64

	given_int := flag.Int64("n", 0, "a number")
	bit_pos := flag.Uint("p", 0, "bit to change position")
	to0 := flag.Bool("to0", false, "change bit to 0")
	to1 := flag.Bool("to1", false, "change bit to 1")

	flag.Parse()

	if *bit_pos > 63 {
		fmt.Println("Max bit position is 63")
		os.Exit(1)
	}

	switch {
	case *to1:
		changed_int = *given_int | (1 << *bit_pos)
	case *to0:
		changed_int = *given_int & ^(1 << *bit_pos)
	}

	fmt.Println(changed_int)
}
