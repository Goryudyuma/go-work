package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scan(&X)
	fmt.Println(X/10 + X%10)
}
