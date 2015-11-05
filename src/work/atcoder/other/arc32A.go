package main

import (
	"fmt"
)

func main() {
	var prime [500501]bool
	prime[0] = true
	prime[1] = true
	for i := 2; i < 500501; i++ {
		if prime[i] == false {
			for j := 2; i*j < 500501; j++ {
				prime[i*j] = true
			}
		}
	}
	var N int
	fmt.Scan(&N)
	N = N * (N + 1) / 2
	if prime[N] == false {
		fmt.Println("WANWAN")
	} else {
		fmt.Println("BOWWOW")
	}
}
