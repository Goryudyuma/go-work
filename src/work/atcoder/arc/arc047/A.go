package main

import (
	"fmt"
)

func main() {
	var N, L int64
	var S string
	fmt.Scan(&N)
	fmt.Scan(&L)
	fmt.Scan(&S)
	ans, nowtabs := int64(0), int64(0)

	for _, i := range S {
		if i == '+' {
			nowtabs++
		} else if i == '-' {
			nowtabs--
		}
		if nowtabs == L {
			nowtabs = 0
			ans++
		}
	}

	fmt.Println(ans)
}
