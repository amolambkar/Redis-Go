package main

import (
	"fmt"
	// Import the redigo/redis package.
	"github.com/gomodule/redigo/redis"
)

func hashset() {
	fmt.Println("Setting Hash Values")
	// Establish a connection to the Redis server listening on port
	// 6379 of the local machine. 6379 is the default port, so unless
	// you've already changed the Redis configuration file this should
	// work.

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err")

	}

	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conn.Close()

	// Send our command across the connection. The first parameter to
	// Do() is always the name of the Redis command (in this example
	// HMSET), optionally followed by any necessary arguments (in this
	// example the key, followed by the various hash fields and values).

	_, err = conn.Do("HMSET", "stu1", "name", "amol", "age", "20", "class", "T.Y")

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Hash Value Set")

	}

}
