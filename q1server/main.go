package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/adjust/rmq"
)

const (
	unackedLimit = 1000
	numConsumers = 2
	batchSize    = 12
)

func main() {
	connection := rmq.OpenConnection("consumer", "tcp", "localhost:6379", 1)
	taskQueue := connection.OpenQueue("q1")
	taskConsumer := &TaskConsumer{}
	taskQueue.AddConsumer("task consumer", taskConsumer)
	select {}
}

type TaskConsumer struct {
	name   string
	count  int
	before time.Time
}

func (consumer *TaskConsumer) Consume(delivery rmq.Delivery) {
	var task Task
	if err = json.Unmarshal([]byte(delivery.Payload()), &task); err != nil {
		// handle error
		delivery.Reject()
		return
	}

	// perform task
	log.Printf("performing task %s", task)
	delivery.Ack()
}
func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	consumer.count++
	if consumer.count%batchSize == 0 {
		duration := time.Now().Sub(consumer.before)
		consumer.before = time.Now()
		perSecond := time.Second / (duration / batchSize)
		log.Printf("%s consumed %d %s %d", consumer.name, consumer.count, delivery.Payload(), perSecond)
	}
	time.Sleep(time.Millisecond)
	if consumer.count%batchSize == 0 {
		delivery.Reject()
	} else {
		delivery.Ack()
	}
}
