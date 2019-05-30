package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var Log *log.Logger

func init() {
	var logpath = "yes-transaction-log"

	flag.Parse()
	FileName, err1 := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(FileName, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
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
	go Q1(redisdb)
	// go Q2(redisdb)
	fmt.Println("before block")
	select {}
}

//Q1 read data from Q1...
func Q1(redisDB *redis.Client) {
	for {
		result, error := redisDB.BLPop(1*time.Second, "payout").Result()
		if error != nil {
			fmt.Println(error.Error())
		}
		Log.Println("===>>> ", "result[0] ", result[0], " Payload ", result[1])
	}
}

//Q2 read data from Q2
func Q2(redisDB *redis.Client) {
	// for {
	result, error := redisDB.BLPop(0, "q2").Result()
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println(result[0], result[1])
	// }
}
