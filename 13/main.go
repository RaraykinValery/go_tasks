package main

import "fmt"

func main() {
	a := 1
	b := 2

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("Swap")

	a, b = b, a

	fmt.Println("a=", a)
	fmt.Println("b=", b)

}
