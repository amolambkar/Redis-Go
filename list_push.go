package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func push() {
	fmt.Println("Push Element to List")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	_, err = conn.Do("LPUSH", "num", "1", "2", "3")

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("List Values Push")

	}
}
