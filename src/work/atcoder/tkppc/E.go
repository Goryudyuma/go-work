//TLE
//WFして一番遠いとこ求めたけど無理っぽい
//最遠点対求めて各点のうち遠い方を出力？
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	WF := make([][]int64, N)
	for i := 0; i < N; i++ {
		WF[i] = make([]int64, N)
		for j := 0; j < N; j++ {
			WF[i][j] = math.MaxInt64 / 6
		}
		WF[i][i] = 0
	}
	for i := 0; i < N-1; i++ {
		a, b, cost := nextInt(), nextInt(), nextInt64()
		a--
		b--
		WF[a][b] = cost
		WF[b][a] = cost
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				maybe := WF[i][k] + WF[k][j]
				if WF[i][j] > maybe {
					WF[i][j] = maybe
					WF[j][i] = maybe
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		ans, memo := WF[i][0], 0
		for j := 1; j < N; j++ {
			if i != j && ans < WF[i][j] {
				ans = WF[i][j]
				memo = j
			}
		}
		fmt.Println(memo + 1)
	}
}
