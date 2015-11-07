package main

import (
	"fmt"
)

func main() {
	var A, B, C, K, S, T int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&K)
	fmt.Scan(&S)
	fmt.Scan(&T)

	ans := 0
	ans = A*S + B*T
	if S+T >= K {
		ans -= C * (S + T)
	}
	fmt.Println(ans)
}
