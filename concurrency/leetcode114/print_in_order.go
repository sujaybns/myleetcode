package leetcode114

import (
	"fmt"
	"sync"
)

type InOrderFooBar struct {
	firstChan  chan int
	secondChan chan int
	thirdChan  chan int
	mysync     sync.WaitGroup
}

func (f *InOrderFooBar) first() {
	defer f.mysync.Done()
	<-f.firstChan
	fmt.Println("first")
	f.secondChan <- 1
}

func (f *InOrderFooBar) second() {
	defer f.mysync.Done()
	<-f.secondChan
	fmt.Println("second")
	f.thirdChan <- 1
}

func (f *InOrderFooBar) third() {
	defer f.mysync.Done()
	<-f.thirdChan
	fmt.Println("third")
	f.firstChan <- 1
}

func InitInOrderFooBar(n int) *InOrderFooBar {
	fooBar := &InOrderFooBar{firstChan: make(chan int, 1),
		secondChan: make(chan int, 1),
		thirdChan:  make(chan int, 1)}
	fooBar.firstChan <- 1
	fooBar.mysync.Add(n)
	return fooBar
}
