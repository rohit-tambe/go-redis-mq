package main

import (
	"fmt"
	"log"

	"github.com/adjust/rmq"
)

const (
	// numDeliveries = 100000000
	numDeliveries = 10
	batchSize     = 10000
)

func main() {
	connection := rmq.OpenConnection("producer", "tcp", "localhost:6379", 2)
	things := connection.OpenQueue("things")
	balls := connection.OpenQueue("balls")

	for i := 0; i < numDeliveries; i++ {

		delivery := fmt.Sprintf("delivery %d %v", i, "rohit")
		ballDelivery := fmt.Sprintf("delivery %d %v", i, "balls")
		log.Printf(delivery)
		log.Printf(ballDelivery)
		things.Publish(delivery)
		balls.Publish(ballDelivery)
		log.Printf("Publish sucessfully ........")
	}
}
