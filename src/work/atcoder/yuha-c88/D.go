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

var data [][]bool
var N int
var p_r []string
var p map[string]int
var max_r int

func ans(S, E, RS string, R []int, f bool) ([]int, bool) {
	if S == E {
		if f {
			if len(R) < max_r {
				max_r = len(R)
			}
			return R, true
		} else {
			return ans(E, RS, RS, R, true)
		}
	} else {
		ret := make([]int, 31)
		if len(R) > max_r {
			return ret, false
		}
		for i := 0; i < N; i++ {
			if data[p[S]][i] {
				data[p[S]][i] = false
				data[i][p[S]] = false
				new_R := append(R, i)
				memo, flag := ans(p_r[i], E, RS, new_R, f)
				if flag {
					if len(ret) > len(memo) {
						ret = memo
					} else if len(ret) == len(memo) {
						var memo1, memo2 string
						for _, x := range ret {
							memo1 += p_r[x]
						}
						for _, x := range memo {
							memo2 += p_r[x]
						}
						if memo1 > memo2 {
							ret = memo
						}
					}
				}

				data[p[S]][i] = true
				data[i][p[S]] = true
			}
		}
		if len(ret) > 30 {
			return ret, false
		} else {
			return ret, true
		}
	}
}

func main() {
	max_r = 30
	N = nextInt()
	p = make(map[string]int, N)
	p_r = make([]string, N)
	for i := 0; i < N; i++ {
		memo := nextString()
		p[memo] = i
		p_r[i] = memo
	}
	data = make([][]bool, N)
	for i := 0; i < N; i++ {
		data[i] = make([]bool, N)
		for j := 0; j < N; j++ {
			data[i][j] = false
		}
	}
	end := nextInt()
	for i := 0; i < end; i++ {
		S, E := nextString(), nextString()
		data[p[S]][p[E]] = true
		data[p[E]][p[S]] = true
	}
	S, G := nextString(), nextString()
	ANS, _ := ans(S, G, S, make([]int, 0), false)
	ANS = append(ANS[len(ANS)-1:], ANS[:len(ANS)-1]...)
	for _, x := range ANS {
		fmt.Printf("%v", p_r[x])
	}
	fmt.Printf("\n")
}
