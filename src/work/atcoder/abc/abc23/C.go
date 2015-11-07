package main

import (
	"bufio"
	"fmt"
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

type point struct {
	x, y int
}

func main() {
	R, C, K := nextInt(), nextInt(), nextInt()
	N := nextInt()
	dataX := make([]int, R)
	dataY := make([]int, C)
	dataZ := make([]point, N)
	for i := 0; i < N; i++ {
		r, c := nextInt(), nextInt()
		r--
		c--
		dataZ[i] = point{x: r, y: c}
		dataX[r]++
		dataY[c]++
	}
	dataX2 := make([]int, K+1)
	dataY2 := make([]int, K+1)
	for i := 0; i < R; i++ {
		if dataX[i] < K+1 {
			dataX2[dataX[i]]++
		}
	}
	for j := 0; j < C; j++ {
		if dataY[j] < K+1 {
			dataY2[dataY[j]]++
		}
	}
	ans := 0
	for k := 0; k < K+1; k++ {
		ans += dataX2[k] * dataY2[K-k]
	}
	for i := 0; i < N; i++ {
		X := dataZ[i].x
		Y := dataZ[i].y
		if dataX[X]+dataY[Y] == K {
			ans--
		}
		if dataX[X]+dataY[Y] == K+1 {
			ans++
		}
	}
	fmt.Println(ans)
}
