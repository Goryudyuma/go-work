package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	i, e := strconv.Atoi(nextString())
	if e != nil {
		panic(e)
	}
	return i
}

func nextInt64() int64 {
	i, e := strconv.ParseInt(nextString(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i

}

func nextFloat() float64 {
	f, e := strconv.ParseFloat(nextString(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func nextString() string {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	return sc.Text()
}

func main() {
	N := nextInt()
	D := make([]int, N)
	for i := 0; i < N; i++ {
		D[i] = nextInt()
	}
	sort.Ints(D)
	DP := make([][]int64, 4)
	for i := 0; i < 4; i++ {
		DP[i] = make([]int64, N)
		if i == 0 {
			for j := 0; j < N; j++ {
				DP[i][j] = int64(1)
			}
		}
	}
	for i := 0; i < 3; i++ {
		m, n := 0, 0
		for n = 0; n < N; n++ {
			for m < N && D[n]*2 > D[m] {
				m++
			}
			if m == N {
				continue
			}
			DP[i+1][m] += DP[i][n]
			DP[i+1][m] %= 1e9 + 7
		}
		for n = 1; n < N; n++ {
			DP[i+1][n] += DP[i+1][n-1]
			DP[i+1][n] %= 1e9 + 7
		}
	}
	ans := int64(0)
	for j := 0; j < N; j++ {
		ans += DP[3][j]
		ans %= 1e9 + 7
	}
	fmt.Println(ans)
}
