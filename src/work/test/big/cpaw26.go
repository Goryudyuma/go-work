package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(32134)
	b := big.NewInt(193127)
	for a.Cmp(b) != 0 {
		if a.Cmp(b) < 0 {
			a = a.Add(a, big.NewInt(1584891))
		} else {
			b = b.Add(b, big.NewInt(3438478))
		}
	}
	fmt.Println(a)
}
