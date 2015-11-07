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
	n := nextInt()
	k := nextInt()
	k++
	data := make([][]int, n+k)
	for i := 0; i < n+k; i++ {
		data[i] = make([]int, (i+1)*2)
		if i != 0 {
			for j := 1; j < i*2; j++ {
				data[i][j] = (data[i-1][j] + data[i-1][j-1]) % (1e9 + 7)
			}
		} else {
			data[0][1] = 1
		}
	}
	fmt.Println(data[n-1+k][n])
}
