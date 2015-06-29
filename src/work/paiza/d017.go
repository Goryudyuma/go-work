package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

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
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	return sc.Text()
}

func main() {
	var l float64
	ans1 := math.Inf(0)
	ans2 := math.Inf(-1)
	for i := 0; i < 5; i++ {
		l = float64(nextInt())
		ans1 = math.Min(ans1, l)
		ans2 = math.Max(ans2, l)
	}
	fmt.Println(ans2)
	fmt.Println(ans1)

}
