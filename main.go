package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	fmt.Println("Welcome to redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "amol", 0).Err() // we can call set with a `Key` and a `Value`.

	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--------------------------------------")
	setvalue()
	fmt.Println("--------------------------------------")
	getvalue()
	fmt.Println("--------------------------------------")
	hashset()
	fmt.Println("--------------------------------------")
	gethash()
	fmt.Println("")
	fmt.Println("--------------------------------------")
	push()
	fmt.Println("--------------------------------------")
	pop()
	fmt.Println("--------------------------------------")
	addset()
	fmt.Println("--------------------------------------")
	countset()
	fmt.Println("--------------------------------------")
	addsortedset()
	fmt.Println("--------------------------------------")
	sortedsetcount()
}
