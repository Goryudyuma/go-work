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
	N, R, C := nextInt(), nextInt(), nextInt()
	data := make([][][]int64, N)
	data_map := make([][][]byte, N)
	for i := 0; i < N; i++ {
		data[i] = make([][]int64, R)
		data_map[i] = make([][]byte, R)
		for j := 0; j < R; j++ {
			data[i][j] = make([]int64, C)
			data_map[i][j] = make([]byte, C)
			data_map[i][j] = []byte(nextString())
			for k := 0; k < C; k++ {
				data[i][j][k] = math.MaxInt64 / 6
			}
		}
	}
	data[0][0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j < R; j++ {
			for k := 0; k < C; k++ {
				if j == 0 && k == 0 {
					continue
				} else if j == 0 {
					if data_map[i][j][k] == 'H' {
						if data[i][j][k-1] < data[i][j][k] {
							data[i+1][j][k] = data[i][j][k-1]
						} else {
							data[i+1][j][k] = data[i][j][k]
						}
						data[i][j][k] = math.MaxInt64 / 6
					} else {
						now, _ := strconv.Atoi(string(data_map[i][j][k]))
						if data[i][j][k] > data[i][j][k-1] {
							data[i][j][k] = data[i][j][k-1]
						}
						data[i][j][k] += int64(now)
					}
				} else if k == 0 {
					if data_map[i][j][k] == 'H' {
						if data[i][j-1][k] < data[i][j][k] {
							data[i+1][j][k] = data[i][j-1][k]
						} else {
							data[i+1][j][k] = data[i][j][k]
						}
						data[i][j][k] = math.MaxInt64 / 6
					} else {
						now, _ := strconv.Atoi(string(data_map[i][j][k]))
						if data[i][j][k] > data[i][j-1][k] {
							data[i][j][k] = data[i][j-1][k]
						}
						data[i][j][k] += int64(now)
					}
				} else {
					if data_map[i][j][k] == 'H' {
						min := data[i][j][k]
						if min > data[i][j-1][k] {
							min = data[i][j-1][k]
						}
						if min > data[i][j][k-1] {
							min = data[i][j][k-1]
						}
						data[i+1][j][k] = min
						data[i][j][k] = math.MaxInt64 / 6
					} else {
						now, _ := strconv.Atoi(string(data_map[i][j][k]))
						min := data[i][j][k]
						if min > data[i][j-1][k] {
							min = data[i][j-1][k]
						}
						if min > data[i][j][k-1] {
							min = data[i][j][k-1]
						}
						data[i][j][k] = min + int64(now)
					}
				}
			}
		}
	}
	fmt.Println(data[N-1][R-1][C-1])
}
