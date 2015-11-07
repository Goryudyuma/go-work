package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, k int
	fmt.Scan(&x, &y, &k)
	fmt.Println(x + y - int(math.Abs(float64(y-k))))
}
