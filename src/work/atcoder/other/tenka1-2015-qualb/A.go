package main

import (
	"fmt"
)

func main() {
	A := make([]int64, 20)
	A[0], A[1], A[2] = 100, 100, 200
	for i := 3; i < 20; i++ {
		A[i] = A[i-1] + A[i-2] + A[i-3]
	}
	fmt.Println(A[19])
}
