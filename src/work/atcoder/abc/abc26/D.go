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

var A, B, C float64

func main() {
	A = float64(nextInt())
	B = float64(nextInt())
	C = float64(nextInt())
	max := float64(1 << 30)
	min := float64(-1)

	mid := (max + min) / 2.0
	for i := 0; i < 1000; i++ {

		mid = (max + min) / 2.0
		if calc(mid) > 100 {
			max = mid
		} else {
			min = mid
		}
	}
	fmt.Println(max)
}

func calc(t float64) (ret float64) {
	ret = A*t + B*math.Sin(C*t*math.Pi)
	return
}
