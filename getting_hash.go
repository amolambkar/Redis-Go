package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func gethash() {
	fmt.Println("getting hash values")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println(err)

	}

	defer conn.Close()

	// Issue a HGET command to retrieve the title for a specific album,
	// and use the Str() helper method to convert the reply to a string.

	name, err := redis.String(conn.Do("HGET", "stu1", "name"))

	if err != nil {
		fmt.Println(err)

	}

	age, err := redis.Int(conn.Do("HGET", "stu1", "age"))

	if err != nil {
		fmt.Println(err)

	}
	class, err := redis.String(conn.Do("HGET", "stu1", "class"))

	if err != nil {
		fmt.Println(err)

	}

	fmt.Printf("%s is of age %d and is in %s class", name, age, class)

}
