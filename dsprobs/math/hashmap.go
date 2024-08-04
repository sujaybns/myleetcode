package math

import "fmt"

const size int = 10

type HashTable struct {
	chain [size]*LinkedList
}

type LinkedList struct {
	head *node
}

type node struct {
	val  string
	next *node
}

func (l *LinkedList) insert(inputString string) {
	mynode := &node{val: inputString}
	second := l.head
	mynode.next = second
	l.head = mynode
}

func (l *LinkedList) search(inputString string) bool {
	temp := l.head
	for temp != nil {
		if temp.val == inputString {
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *LinkedList) delete(inputString string) string {
	if !l.search(inputString) {
		return inputString + "not found"
	}
	if l.head.val == inputString {
		l.head = l.head.next
	} else {
		temp := l.head
		for temp.next.val != inputString {
			temp = temp.next
		}
		temp.next = temp.next.next
	}
	return ""
}

func (h *HashTable) Insert(inputString string) {
	index := hash(inputString)
	h.chain[index].insert(inputString)
}

func (h *HashTable) Delete(inputString string) {
	index := hash(inputString)
	msg := h.chain[index].delete(inputString)
	if msg != "" {
		fmt.Println(msg)
	}
}

func (h *HashTable) Search(inputString string) bool {
	index := hash(inputString)
	return h.chain[index].search(inputString)
}

func hash(input string) int {
	sum := 0
	for _, val := range input {
		sum = sum + int(val)
	}
	return sum % size
}

func Init() *HashTable {
	hashTable := HashTable{}
	for i := 0; i < size; i++ {
		hashTable.chain[i] = &LinkedList{}
	}
	return &hashTable
}
