package main

import (
	"fmt"
)

func main() {
	var N, A int
	M := make(map[int]int)
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		M[A] += 1
	}
	fmt.Println(N - len(M))
}
