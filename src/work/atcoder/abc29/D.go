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
	var N,
		ans,
		M,
		scale int64
	N = nextInt64()
	N++
	ans = int64(0)
	M = N
	scale = int64(1)

	for M != int64(0) {
		now := M % 10
		if now > 1 {
			ans += scale
		} else if now == 1 {
			ans += N % scale
		}
		ans += M / 10 * scale
		scale *= 10
		M /= 10
	}
	fmt.Println(ans)
}
