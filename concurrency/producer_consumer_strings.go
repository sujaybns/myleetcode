package concurency

import (
	"fmt"
	"time"
)

type ProducerConsumerOfStrings struct {
	bufferedChannel chan string
	exit            chan string
}

// Producer function generates strings and sends them to the channel
func (p *ProducerConsumerOfStrings) producer(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Message %d", i)
		fmt.Println("Produced:", msg)
		ch <- msg
		time.Sleep(time.Millisecond * 500) // Simulate some delay in production
	}
	close(ch) // Close the channel when done producing
}

// Consumer function receives strings from the channel and processes them
func (p *ProducerConsumerOfStrings) consumer(ch <-chan string, exit bool) {
	for msg := range ch {
		fmt.Println("Consumed:", msg)
		time.Sleep(time.Second) // Simulate some delay in consumption
	}
	if exit {
		close(p.exit)
	}
}

func InitProducerConsumerStrings() *ProducerConsumerOfStrings {
	// Create a buffered channel with a buffer size of 2
	pcStrings := &ProducerConsumerOfStrings{bufferedChannel: make(chan string, 2), exit: make(chan string)}
	return pcStrings
}
