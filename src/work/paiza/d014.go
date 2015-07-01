package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	for i := 0; i < len(S); i++ {
		fmt.Printf("%c", S[i]-'a'+'A')
	}
	fmt.Printf("\n")
}
