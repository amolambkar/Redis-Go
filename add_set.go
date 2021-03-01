package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func addset() {
	fmt.Println("Add Value to set")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	_, err = conn.Do("SADD", "set", "1", "2", "3")

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("SET elements Added")

	}
}
