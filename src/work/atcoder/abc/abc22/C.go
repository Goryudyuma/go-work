package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M int64
	fmt.Scan(&N, &M)
	data := make([][]int64, N)
	for i := int64(0); i < N; i++ {
		data[i] = make([]int64, N)
		for j := int64(0); j < N; j++ {
			data[i][j] = math.MaxInt64 / 6
		}
		data[i][i] = 0
	}
	var u, v, l int64
	for i := int64(0); i < M; i++ {
		fmt.Scanln(&u, &v, &l)
		u--
		v--
		if data[u][v] > l {
			data[u][v] = l
			data[v][u] = data[u][v]
		}
	}

	//WF
	for k := int64(1); k < N; k++ {
		for i := int64(1); i < N; i++ {
			for j := int64(1); j < N; j++ {
				if data[i][j] > data[i][k]+data[k][j] {
					data[i][j] = data[i][k] + data[k][j]
					data[j][i] = data[i][j]
				}
			}
		}
	}

	ans := int64(math.MaxInt64 / 7)
	for start := int64(1); start < N; start++ {
		for end := start + int64(1); end < N; end++ {
			if ans > data[0][start]+data[start][end]+data[end][0] {
				ans = data[0][start] + data[start][end] + data[end][0]
			}
		}
	}
	if ans >= int64((math.MaxInt64 / 7)) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
