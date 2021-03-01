package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func setvalue() {

	fmt.Println("Setting Values")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	err := client.Set("name1", "amol", 0).Err() // we can call set with a `Key` and a `Value`.

	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value set")
	}
}
