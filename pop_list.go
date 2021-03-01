package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func pop() {
	fmt.Println("POP Element From List")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	popitem, err := conn.Do("LPOP", "num")

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("List Value POP")
		fmt.Printf("%s is pop fromm list", popitem)
		fmt.Println("")
	}

}
