package main

import (
	"fmt"
)

func main() {
	var N int
	var A int
	var B int
	var S string
	var d int

	now := 0

	fmt.Scan(&N)
	fmt.Scan(&A)
	fmt.Scan(&B)

	for i := 0; i < N; i++ {
		fmt.Scan(&S)
		fmt.Scan(&d)
		if S == "East" {
			if d < A {
				now += A
			} else if d > B {
				now += B
			} else {
				now += d
			}
		} else {
			if d < A {
				now -= A
			} else if d > B {
				now -= B
			} else {
				now -= d
			}
		}
	}
	if now == 0 {
		fmt.Println(0)
	} else if now > 0 {
		fmt.Printf("East %v\n", now)
	} else {
		fmt.Printf("West %v\n", -now)
	}
}
