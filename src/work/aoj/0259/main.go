package main
 
import (
    "fmt"
    "sort"
    "strconv"
)
 
var memo []int
 
func solve(v int) int {
    if v == 6174 {
        return 0
    }
    if memo[v] == 0 {
        a := make([]int, 4)
        for i := 0; i < 4; i++ {
            a[i] = v % 10
            v /= 10
        }
        sort.Slice(a, func(i, j int) bool {
            return a[i] > a[j]
        })
        l := int(0)
        for i := 0; i < 4; i++ {
            l *= 10
            l += a[i]
        }
        sort.Slice(a, func(i, j int) bool {
            return a[i] < a[j]
        })
        s := int(0)
        for i := 0; i < 4; i++ {
            s *= 10
            s += a[i]
        }
        memo[v] = solve(l-s) + 1
    }
    return memo[v]
}
 
func main() {
    var v int
    memo = make([]int, 10000)
    for {
        var s string
        fmt.Scan(&s)
        v, _ = strconv.Atoi(s)
        if v == 0 {
            break
        }
        if v%1111 == 0 {
            fmt.Println("NA")
            continue
        }
        fmt.Println(solve(v))
    }
}
