package main

import (
	"fmt"
)

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		{
			s += a[i]
		}
	}
	return s
}

func main() {
	var x = []int{2, 3, 4}
	fmt.Println(sum(x))
}
