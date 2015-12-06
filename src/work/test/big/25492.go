package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N int
	fmt.Scan(&N)
	D1, D2 := make([]*big.Int, 4), make([]*big.Int, 4)
	for i := 0; i < 4; i++ {
		D1[i] = big.NewInt(0)
		D2[i] = big.NewInt(0)
	}
	D1[1] = big.NewInt(1)
	ans := big.NewInt(0)

	for i := 0; i < N; i += 2 {
		D2[1].Add(D1[2], big.NewInt(0))
		D2[2].Add(D1[2], big.NewInt(0))
		D2[2].Add(D2[2], D1[1])
		D1[1].Mul(D1[1], big.NewInt(2))
		D2[1].Add(D2[1], D1[1])
		D1[1].Add(D2[1], big.NewInt(0))
		D1[2].Add(D2[2], big.NewInt(0))
		ans.Add(ans, D1[1])
	}

	fmt.Println(ans)

}
