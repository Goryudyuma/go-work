package main

import (
	"fmt"
)

func main() {
	var x []int = []int{1, 2, 3}
	s := Sum(x)
	fmt.Println(s)
}

func Sum(array []int) int {
	ret := 0
	for i := 0; i < len(array); i++ {
		ret += array[i]
	}
	return ret
}
