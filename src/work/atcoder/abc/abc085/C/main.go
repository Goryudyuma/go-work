package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if i*10000+j*5000+(n-i-j)*1000 == y {
				fmt.Println(i, j, n-i-j)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
