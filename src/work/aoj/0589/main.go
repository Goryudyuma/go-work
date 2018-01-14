package main

import (
	"fmt"
	"sort"
)

type node struct {
	Key string
	Num int
}

func main() {
	var n int
	fmt.Scan(&n)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		var num int
		fmt.Scan(&s, &num)
		if _, ok := m[s]; !ok {
			m[s] = 0
		}
		m[s] += num
	}
	list := make([]node, 0, len(m))
	for k, v := range m {
		list = append(list, node{Key: k, Num: v})
	}
	sort.Slice(list, func(i, j int) bool {
		if len(list[i].Key) == len(list[j].Key) {
			return list[i].Key < list[j].Key
		} else {
			return len(list[i].Key) < len(list[j].Key)
		}
	})
	for _, v := range list {
		fmt.Println(v.Key, v.Num)
	}
}
