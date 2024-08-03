package math

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Linkedist struct {
	head   *Node
	length int
}

func (l *Linkedist) Insert(newNode *Node) {
	second := l.head
	newNode.next = second
	l.head = newNode
	l.length++
}

func (l Linkedist) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}

func (l Linkedist) search(val int) bool {
	temp := l.head
	for temp != nil {
		if temp.val == val {
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *Linkedist) Delete(newNode *Node) {
	if !l.search(newNode.val) {
		fmt.Println("Number not found")
		return
	}
	if l.head.val == newNode.val {
		l.head = l.head.next
	} else {
		temp := l.head
		for temp.next.val != newNode.val {
			temp = temp.next
		}
		temp.next = temp.next.next
	}
	l.length--
}
