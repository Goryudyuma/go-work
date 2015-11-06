package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Point struct {
	X, Y float64
	Name string
}

func main() {
	var network bytes.Buffer
	P, Q := Point{X: 10, Y: 20, Name: "Point P"}, Point{X: 30, Y: 10, Name: "Point Q"}
	enc := gob.NewEncoder(&network)
	enc.Encode(P)
	fmt.Println(P)
	fmt.Println(Q)
	dec := gob.NewDecoder(&network)
	var R Point
	dec.Decode(&R)
	fmt.Println(R)
}
