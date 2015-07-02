package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)

	fmt.Printf("%v %v\n", m/n, m%n)
}
