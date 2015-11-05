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
	var G, C, P float64
	fmt.Scan(&G)
	fmt.Scan(&C)
	fmt.Scan(&P)
	S := nextString()
	var G2, C2, P2 float64
	for i := 0; i < len(S); i++ {
		if S[i] == 'G' {
			G2++
		} else if S[i] == 'C' {
			C2++
		} else {
			P2++
		}
	}
	ans := 0.0
	x := math.Min(G, C2)
	G -= x
	C2 -= x

	ans += x * 3

	x = math.Min(C, P2)
	C -= x
	P2 -= x

	ans += x * 3

	x = math.Min(P, G2)
	P -= x
	G2 -= x

	ans += x * 3

	x = math.Min(G, G2)
	G -= x
	G2 -= x

	ans += x

	x = math.Min(C, C2)
	C -= x
	C2 -= x

	ans += x

	x = math.Min(P, P2)
	P -= x
	P2 -= x

	ans += x

	fmt.Println(int(ans))
}
