package concurency

import (
	"fmt"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	prodcons := initProducerConsumer()
	go prodcons.producer()
	go prodcons.consumer()
	<-prodcons.exit
	fmt.Println("Producer consumer completed")
}

func TestProducerConsumerForSelectCase(t *testing.T) {
	prodcons := initProducerConsumer()
	go prodcons.producer()
	go prodcons.consumerForSelectClose()
	<-prodcons.exit
	fmt.Println("Producer consumer completed")
}
