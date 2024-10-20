package leetcode1115

import ("fmt"
		"sync"
)

type FooBar struct {
	n          int
	myOddchan  chan struct{}
	myEvenchan chan struct{}
	wg         sync.WaitGroup
}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n,
		myOddchan:  make(chan struct{}, 1),
		myEvenchan: make(chan struct{}, 1),
	}
}

func printFooBarAlternate() {
	fooBar := NewFooBar(10)
	fooBar.wg.Add(2)
	go fooBar.printFoo()
	go fooBar.printBar()
	fooBar.wg.Wait() 
}

func (fb *FooBar) printFoo() {
	defer fb.wg.Done()
	for i := 0; i < fb.n; i = i + 2 {
		fb.myOddchan <- struct{}{} //Enable Fooc- print Foo
		fmt.Println(i)
		fb.myEvenchan <- struct{}{} //Block Foo
	}
}

func (fb *FooBar) printBar() {
	defer fb.wg.Done()
	for i := 1; i < fb.n; i = i + 2 {
		<-fb.myEvenchan //Block Bar
		fmt.Println(i)
		<-fb.myOddchan //Enable Fooc- print Foo
	}
}