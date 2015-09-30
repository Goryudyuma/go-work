package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Do("SELECT", "1")
	c.Do("SET", "abc", "test")
	c.Do("SET", "def", "test23")
	str, err := redis.String(c.Do("GET", "abc"))
	if err != nil {
		fmt.Println("key not found")
	}
	fmt.Println(str)
}
