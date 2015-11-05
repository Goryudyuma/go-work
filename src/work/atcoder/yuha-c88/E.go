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
	data := make([]string, N)
	for i := 0; i < N; i++ {
		data[i] = nextString()
	}
	D4 := make([]string, 4)
	YA := make([]string, 4)
	D4[0] = nextString()
	YA[0] = nextString()
	str := nextString()
	D4[1] = string([]rune(str)[:1])
	YA[1] = string([]rune(str)[1:2])
	YA[2] = string([]rune(str)[3:4])
	D4[2] = string([]rune(str)[4:5])
	YA[3] = nextString()
	D4[3] = nextString()
	var ans, memo string = "", ""
	for i := 0; i < N; i++ {
		if YA[0] == "↓" {
			if string([]rune(data[i])[:1]) == D4[0] {
				ans = string([]rune(data[i])[1:])
			}
		}
		if YA[0] == "↑" {
			if string([]rune(data[i])[1:]) == D4[0] {
				ans = string([]rune(data[i])[:1])
			}
		}
		if ans != memo {
			memo = ans
			check := make([]bool, 3)
			check_word := make([]string, 3)
			if YA[1] == "→" {
				check_word[0] = D4[1] + ans
			} else {
				check_word[0] = ans + D4[1]
			}
			if YA[2] == "→" {
				check_word[1] = ans + D4[2]
			} else {
				check_word[1] = D4[2] + ans
			}
			if YA[3] == "↓" {
				check_word[2] = ans + D4[3]
			} else {
				check_word[2] = D4[3] + ans
			}
			for j := 0; j < N; j++ {
				for k := 0; k < 3; k++ {
					if check_word[k] == data[j] {
						check[k] = true
					}
				}
			}
			if check[0] && check[1] && check[2] {
				fmt.Println(ans)
				return
			}
		}

	}
}
