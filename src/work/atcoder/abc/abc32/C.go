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

func main() {
	N, K := nextInt64(), nextInt64()
	S := make([]int64, N)
	for i := int64(0); i < N; i++ {
		S[i] = nextInt64()
		if S[i] == 0 {
			fmt.Println(N)
			return
		}
	}
	x, y, mul, ans := int64(0), int64(0), int64(1), int64(0)
	for y < N {
		if mul <= K || x == y {
			mul *= S[y]
			y++
		} else {
			mul /= S[x]
			x++
		}
		if mul <= K && ans < y-x {
			ans = y - x
		}
	}
	fmt.Println(ans)
}
