package main

import (
	"fmt"
)

func main() {
	var N, D, K int
	fmt.Scan(&N, &D, &K)

	var data [][2]int = make([][2]int, D)
	for i := 0; i < D; i++ {
		fmt.Scan(&data[i][0], &data[i][1])
	}
	for i := 0; i < K; i++ {
		var S, T, nowx, nowy int
		fmt.Scan(&S, &T)
		nowx = S
		nowy = S
		for j := 0; j < D; j++ {
			if nowx <= data[j][1] && nowx > data[j][0] {
				nowx = data[j][0]
			}
			if nowy >= data[j][0] && nowy < data[j][1] {
				nowy = data[j][1]
			}
			if nowx <= T && T <= nowy {
				fmt.Println(j + 1)
				break
			}
		}
	}
}
