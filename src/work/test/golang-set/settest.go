package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

func main() {
	set := mapset.NewSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add("a")
	set.Add(1)

	fmt.Println(set.Cardinality())
	fmt.Println(set)

	var x []int64
	x = append(x, 12345, 12365, 1238798, 12345, 5647879)
	fmt.Println(x)
	var y []interface{}
	for i := range x {
		y = append(y, x[i])
	}
	fmt.Println(set.Union(mapset.NewSetFromSlice(y)))
}
