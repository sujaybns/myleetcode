package concurency

import (
	"fmt"
	"testing"
)

func TestProducerConsumerStrings_OneGoRoutine(t *testing.T) {
	// prodcons := InitProducerConsumerStrings()
	// Start the producer and consumer
	prodcons := InitProducerConsumerStrings()
	go prodcons.producer(prodcons.bufferedChannel)
	prodcons.consumer(prodcons.bufferedChannel, false)

	fmt.Println("Done.")

}

func TestProducerConsumerStrings_TwoGoRoutine(t *testing.T) {
	// prodcons := InitProducerConsumerStrings()
	// Start the producer and consumer
	prodcons := InitProducerConsumerStrings()
	go prodcons.producer(prodcons.bufferedChannel)
	go prodcons.consumer(prodcons.bufferedChannel, true)
	<-prodcons.exit
	fmt.Println("Done.")

}
