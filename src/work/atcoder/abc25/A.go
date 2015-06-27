package main

import (
	"fmt"
)

func main() {
	var S string
	var N int
	fmt.Scan(&S)
	fmt.Scan(&N)
	fmt.Printf("%s%s\n", (string(S[(N-1)/len(S)])), (string(S[(N-1)%len(S)])))
}
