package main

import (
	"fmt"
)

func main() {
	var N, S, T, W, A int
	fmt.Scan(&N, &S, &T)
	count := 0
	for i := 0; i < N; i++ {
		if i == 0 {
			fmt.Scan(&W)
		} else {
			fmt.Scan(&A)
			W += A
		}
		if S <= W && W <= T {
			count++
		}
	}
	fmt.Println(count)
}
