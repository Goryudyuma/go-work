package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	P := strings.Split(readLine(), " ")
	N_, _ := strconv.Atoi(P[0])
	a_, _ := strconv.Atoi(P[1])
	N := int64(N_)
	a := int64(a_)
	ks := readLine()
	b, c := make([]int64, N+1), make([]int64, N+1)
	P = strings.Split(readLine(), " ")
	for i := int64(1); i <= N; i++ {
		c_, _ := strconv.Atoi(P[i-1])
		c := int64(c_)
		b[i] = c
	}
	x := int64(0)
	x = 1
	c[a] = 1
	k := big.NewInt(0)
	k, _ = k.SetString(ks, 10)
	for k.Sign() != 0 && c[b[a]] == 0 {
		x++
		a = b[a]
		c[a] = x
		k.Sub(k, big.NewInt(1))
	}
	k.Mod(k, big.NewInt(c[a]-c[b[a]]+1))
	for k.Sign() != 0 {
		a = b[a]
		k.Sub(k, big.NewInt(1))
	}
	fmt.Println(a)
}
