package main

import (
	"fmt"

	"github.com/adjust/redismq"
)

func main() {
	testQueue := redismq.CreateQueue("localhost", "6379", "", 9, "q1")
	// for i := 0; i < 10; i++ {
	// 	testQueue.Put("i")
	// }
	// testQueue.Put("i")
	// testQueue.Put("j")
	// testQueue.Put("j")
	// testQueue.Put("l")
	consumer, err := testQueue.AddConsumer("q1")
	if err != nil {
		panic(err)
	}
	// if error != nil {
	// 	fmt.Println(error)
	// }
	for
	// i := 0; i < 20; i++
	{
		p, err := consumer.Get()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p.CreatedAt)
		fmt.Println(p.Payload)

		err = p.Ack()
		if err != nil {
			fmt.Println(err)
		}
	}
}
