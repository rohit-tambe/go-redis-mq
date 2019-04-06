package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(client)
	q1Channel := make(chan string)
	// q2Channel := make(chan string)
	for {
		go Q1Result(q1Channel)
		result1 := <-q1Channel
		fmt.Println("Result from q1 Queue ", result1)
		// go Q2Result(q2Channel)
		// result2 := <-q2Channel
		// fmt.Println("Result from q2 Queue ", result2)
	}

}
func Q1Result(result chan string) {
	q1Result, err := client.BLPop(0, "redismq::q1").Result()
	if err != nil {
		panic(err)
	}
	result <- q1Result[1]
}
func Q2Result(result chan string) {
	q1Result, err := client.BLPop(1, "q2").Result()
	if err != nil {
		panic(err)
	}
	result <- q1Result[1]
}
