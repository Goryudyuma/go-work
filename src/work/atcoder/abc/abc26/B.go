package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	data := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}
	sort.Float64s(data)
	ans := 0.0
	for i := 0; i < N; i++ {
		j := N - i - 1
		if i%2 == 1 {
			ans -= float64(data[j] * data[j])
		} else {
			ans += float64(data[j] * data[j])
		}
	}
	fmt.Println(ans * math.Pi)
}
