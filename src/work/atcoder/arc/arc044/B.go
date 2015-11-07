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

func mul(a, b int64) int64 {
	return a * b % (1e9 + 7)
}

func pow(a, n int64) int64 {
	var ret int64 = 1
	for n > 0 {
		if n%2 == 1 {
			ret = mul(ret, a)
		}
		a = mul(a, a)
		n >>= 1
	}
	return ret
}

func main() {
	N := nextInt()
	D := make([]int64, N)
	max := 0
	ans := int64(1)
	b := nextInt()
	if b != 0 {
		ans = 0
	}
	D[0]++
	for i := 1; i < N; i++ {
		a := nextInt()
		D[a]++
		if max < a {
			max = a
		}
	}
	if D[0] != 1 {
		ans = 0
	}
	for i := 1; i <= max; i++ {
		for j := int64(0); j < D[i]; j++ {
			ans = mul(ans, pow(2, D[i-1])+1e9+6)
		}
		ans = mul(ans, (pow(2, (D[i] * (D[i] - 1) / 2))))
	}
	fmt.Println(ans)
}
