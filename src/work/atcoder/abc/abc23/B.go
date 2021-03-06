package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var N int
	var T string
	S := "b"
	fmt.Scan(&N)
	fmt.Scan(&T)
	if N%2 == 0 {
		fmt.Println("-1")
	} else {
		for i := 1; utf8.RuneCountInString(S) != N; i++ {
			if i%3 == 1 {
				S = "a" + S + "c"
			} else if i%3 == 2 {
				S = "c" + S + "a"
			} else {
				S = "b" + S + "b"
			}
		}
		if T == S {
			fmt.Println(N / 2)
		} else {
			fmt.Println("-1")
		}
	}
}
