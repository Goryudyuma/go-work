package main

import (
	"fmt"
)

func x(v int) int {
	if v == 0 {
		return 0
	}
	if v%10 == 3 {
		return x(v/10) + 1
	} else {
		return x(v / 10)
	}
}

func main() {
	c := 0
	for i := 0; i < 1000000; i++ {
		if x(i) > 2 {
			c++
		}
	}
	fmt.Println(c)
}
