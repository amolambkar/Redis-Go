package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func countset() {
	fmt.Println("POP Element From List")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	count, err := redis.Int(conn.Do("scard", "set"))

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("Number of elements in set are %d ", count)

		fmt.Println("")
	}

}
