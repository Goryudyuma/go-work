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
	count := 0
	now := N
	for now > 0 {
		now = int(now / 2)
		count++
	}
	flag := true
	if count%2 == 0 {
		now = 1
		for now <= N {
			now *= 2
			if flag {
				flag = false
			} else {
				flag = true
				now++
			}
		}
	} else {
		flag = false
		now = 1
		for now <= N {
			now *= 2
			if flag {
				flag = false
			} else {
				flag = true
				now++
			}
		}
	}
	if (flag && count%2 == 1) || !(flag || count%2 == 1) {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Takahashi")
	}
}
