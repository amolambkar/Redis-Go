package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func getvalue() {

	fmt.Println("Getting values")

	client := redis.NewClient(&redis.Options{

		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := client.Get("name").Result()

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Get Value")

	}

	fmt.Println(val)
}
