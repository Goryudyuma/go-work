package main

import (
	"fmt"
)

func main() {
	var N, T, A, ans, S, G int
	ans = 0
	fmt.Scan(&N)
	fmt.Scan(&T)
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		if G <= A {
			ans += G - S
			S = A
			G = A + T
		} else {
			G = A + T
		}
	}
	ans += G - S
	fmt.Println(ans)

}
