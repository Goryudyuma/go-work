package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	primes := make([]int64, 1)
	primes_f := make([]float64, 1)
	primes[0] = 2
	primes_f[0] = 2.0
	var max int64 = 10000

	start := time.Now()
	var n int64 = 3
	for n = 3; n < max; n += 2 {
		flag := true
		f := float64(n)
		rf := math.Sqrt(f)
		for i := 1; i < len(primes); i++ {
			if primes_f[i] > rf {
				break
			} else if (n % primes[i]) == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, n)
			primes_f = append(primes_f, f)
		}
	}
	goal := time.Now()

	fmt.Printf("%v以下の素数:%v\n", max, primes)
	fmt.Printf("%v経過\n", goal.Sub(start))
}
