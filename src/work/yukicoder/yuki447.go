package main

import (
	"fmt"
)

func Gcd(x int64, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return Gcd(y, x%y)
	}
}

func main() {
	var x, y, gcd, ans int64
	fmt.Scan(&ans, &gcd)
	f := 0
	for i := 0; i < 2; i++ {
		fmt.Scan(&x, &y)
		j := int64(0)
		for ; ans == 0 || (ans%y != x && j < y); j++ {
			ans += gcd
		}
		if j == y {
			f = 1
		}
		gcd = gcd * y / Gcd(gcd, y)
	}
	if f == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
