package main

import (
	"fmt"
	"math/big"
)

func main() {
	var bi big.Int

	a := big.NewInt(1 << 50)
	b := big.NewInt(1 << 40)

	c := bi.Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, c)

	c = bi.Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, c)

	c = bi.Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, c)

	c = bi.Div(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, c)
}
