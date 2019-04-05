package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	redisdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// redisdb.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
	// 	return func(cmd redis.Cmder) error {
	// 		fmt.Printf("starting processing: <%s>\n", cmd)
	// 		err := old(cmd)
	// 		fmt.Printf("finished processing: <%s>\n", cmd)
	// 		return err
	// 	}
	// })
	fmt.Println(redisdb)
	// for {
	go Q1(redisdb)
	go Q2(redisdb)
	// if error != nil {
	// 	fmt.Println(error.Error())
	// }
	// }
	fmt.Println("before block")
	select {}
}

// read data from Q1...
func Q1(redisDB *redis.Client) {
	for {
		result, error := redisDB.BLPop(0, "q1").Result()
		if error != nil {
			fmt.Println(error.Error())
		}
		fmt.Println(result[0], result[1])
	}
}

// read data from Q2
func Q2(redisDB *redis.Client) {
	for {
		result, error := redisDB.BLPop(0, "q2").Result()
		if error != nil {
			fmt.Println(error.Error())
		}
		fmt.Println(result[0], result[1])
	}
}
