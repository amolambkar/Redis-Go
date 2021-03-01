package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func addsortedset() {
	fmt.Println("Add Value to sorted set")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	_, err = conn.Do("ZADD", "set1", "1", "a", "2", "c", "3", "b", "4", "e")

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Sorted SET elements Added")

	}
}
