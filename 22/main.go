package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)

	a.SetString("10000000000", 10)
	b.SetString("1000000000", 10)

	c := new(big.Int).Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, c)

	c = new(big.Int).Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, c)

	c = new(big.Int).Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, c)

	c = new(big.Int).Div(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, c)
}
