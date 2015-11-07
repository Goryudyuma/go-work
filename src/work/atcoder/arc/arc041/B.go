package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	N, M := nextInt(), nextInt()
	data := make([][]int, N)
	for i := 0; i < N; i++ {
		data[i] = make([]int, M)
		str := nextString()
		str_d := strings.Split(str, "")
		for j := 0; j < M; j++ {
			data[i][j], _ = strconv.Atoi(str_d[j])
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if i == 0 {
				fmt.Printf("0")
			} else {
				if data[i-1][j] == 0 {
					fmt.Printf("0")
				} else {
					fmt.Printf("%v", data[i-1][j])
					data[i][j-1] -= data[i-1][j]
					data[i][j+1] -= data[i-1][j]
					data[i+1][j] -= data[i-1][j]
				}
			}
		}
		fmt.Println()
	}
}
