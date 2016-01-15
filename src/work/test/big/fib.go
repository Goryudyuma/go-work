package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < 10000; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	fmt.Println(a)
}
