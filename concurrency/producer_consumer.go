package concurency

import (
	"fmt"
	"strconv"
)

type ProducerConsumer struct {
	blockedQueue chan string
	exit         chan string
}

func (p *ProducerConsumer) producer() {
	for i := 0; i < 5; i++ {
		numStr := strconv.Itoa(i)
		str := "message" + numStr
		p.blockedQueue <- str
	}
	close(p.blockedQueue)
}

func (p *ProducerConsumer) consumer() {
	defer close(p.exit)

	for val := range p.blockedQueue {
		fmt.Println(string(val))
	}
}

func (p *ProducerConsumer) consumerForSelectClose() {
	for {
		select {
		case val, ok := <-p.blockedQueue:
			if !ok {
				close(p.exit)
				return
			}
			fmt.Println("Printing val:" + val)
		default:
			fmt.Print("Doing nothing:")
		}
	}
}

func initProducerConsumer() *ProducerConsumer {
	producerConsumer := &ProducerConsumer{blockedQueue: make(chan string, 5), exit: make(chan string, 5)}
	return producerConsumer
}