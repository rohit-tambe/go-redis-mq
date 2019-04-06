package main

import (
	"fmt"
	"log"
	"time"

	"github.com/adjust/rmq"
)

const (
	unackedLimit = 1000
	numConsumers = 10
	batchSize    = 1
)

func main() {
	connection := rmq.OpenConnection("consumer", "tcp", "localhost:6379", 2)

	go ThingsQ(connection)
	go BallQ(connection)
	go GetRejectedEntry(connection)
	fmt.Println("Before select")
	select {}
}
// ThingsQ get payload
func ThingsQ(connection rmq.Connection) {
	queue := connection.OpenQueue("things")
	queue.StartConsuming(unackedLimit, time.Millisecond)
	for i := 0; i < numConsumers; i++ {

		name := fmt.Sprintf("consumer %d", i)
		consumerName := queue.AddConsumer(name, NewConsumer(i))
		// con := NewConsumer(i)
		fmt.Println(consumerName)
	}
}
func GetRejectedEntry(connection rmq.Connection) {
	queue := connection.OpenQueue("things")
	returned := queue.ReturnAllRejected()
	log.Printf("queue returner returned %d rejected deliveries", returned)
}
func BallQ(connection rmq.Connection) {
	queue := connection.OpenQueue("balls")
	queue.StartConsuming(unackedLimit, time.Millisecond)
	for i := 0; i < numConsumers; i++ {

		name := fmt.Sprintf("consumer %d", i)
		consumerName := queue.AddConsumer(name, NewConsumer(i))
		// con := NewConsumer(i)
		fmt.Println(consumerName)
	}
}

type Consumer struct {
	name   string
	count  int
	before time.Time
}

func NewConsumer(tag int) *Consumer {
	return &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
		count:  0,
		before: time.Now(),
	}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	consumer.count++
	if consumer.count%batchSize == 0 {
		duration := time.Now().Sub(consumer.before)
		consumer.before = time.Now()
		perSecond := time.Second / (duration / batchSize)
		log.Println(delivery.Payload(), perSecond)
		// log.Printf("%s consumed %d %s %d", consumer.name, consumer.count, delivery.Payload(), perSecond)
	}
	time.Sleep(time.Millisecond)
	if consumer.count%batchSize == 0 {
		log.Println("In reject")
		delivery.Reject()
	} else {
		log.Println("In ack")
		delivery.Ack()
	}
}
