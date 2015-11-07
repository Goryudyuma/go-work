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

type Salary struct {
	max int
	min int
}

func main() {
	N := nextInt()
	data := make([]Salary, N)
	for i := 0; i < N; i++ {
		data[i].max = -1
		data[i].min = 1 << 60
	}
	boss := make([]int, N)
	for i := 1; i < N; i++ {
		p := nextInt()
		boss[i] = p - 1
	}
	var ans int
	for i := N - 1; i >= 0; i-- {
		var now_salary int
		if data[i].max == -1 {
			data[i].max = 0
			data[i].min = 0
		}
		now_salary = data[i].max + data[i].min + 1
		if i == 0 {
			ans = now_salary
		} else {
			if data[boss[i]].max < now_salary {
				data[boss[i]].max = now_salary
			}
			if data[boss[i]].min > now_salary {
				data[boss[i]].min = now_salary
			}
		}
	}
	fmt.Println(ans)
}
