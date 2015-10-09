package main

import (
	"log"
	"time"
)

func write(n int) {
	log.Print(n)
	for i := 0; i < 100; i++ {
		log.Print(i, n)
	}
}

func main() {
	log.Print("A")
	for i := 0; i < 100; i++ {
		write(i)
		time.Sleep(time.Minute)
	}
}
