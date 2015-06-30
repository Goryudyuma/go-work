package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := N; i > 0; i-- {
		fmt.Println(i)
	}
}
