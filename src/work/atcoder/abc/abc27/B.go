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
	data := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		data[i] = nextInt()
		sum += data[i]
	}
	if sum%N != 0 {
		fmt.Println(-1)
	} else {
		ave := sum / N
		num := 0
		count := 0
		for i := 0; i < N; i++ {
			num += data[i]
			num -= ave
			if num != 0 {
				count++
			}
		}
		fmt.Println(count)
	}

}
