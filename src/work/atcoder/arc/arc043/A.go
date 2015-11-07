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
	N, A, B := nextInt(), nextInt(), nextInt()
	S := make([]int, N)
	min, max := float64(1<<30), float64(-1)
	ave := float64(0)
	for i := 0; i < N; i++ {
		S[i] = nextInt()
		num := float64(S[i])
		ave += num
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}
	sub := max - min
	ave /= float64(N)
	if sub == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(float64(B)/sub, (float64(A) - ave*float64(B)/sub))
	}
}
