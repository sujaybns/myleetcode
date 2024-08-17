package math

import "fmt"

type BFSNode struct {
	val   int
	left  *BFSNode
	right *BFSNode
}

type BFSTree struct {
	root *BFSNode
}

func (b *BFSTree) Insert(val int) {
	b.root = recBFSInsert(b.root, val)
}

func recBFSInsert(root *BFSNode, val int) *BFSNode {
	if root == nil {
		root = &BFSNode{val: val}
	} else if val > root.val {
		root.right = recBFSInsert(root.right, val)
	} else {
		root.left = recBFSInsert(root.left, val)
	}
	return root
}


func (b *BFSTree) Search(val int) bool {
	return recBFSSearch(b.root, val)
}

func recBFSSearch(root *BFSNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.val {
		return true
	}
	if val > root.val {
		return recBFSSearch(root.right, val)
	}
	return recBFSSearch(root.left, val)
}



func (b *BFSTree) BFSTraversal() []interface{}{
	return recBFSTraversal(b.root)
}

func recBFSTraversal(root *BFSNode)  []interface{}{
	queue := Queue{} 
	queue.Enqueue(root)
	var newList []interface{}
	for !queue.IsEmpty(){
		size := queue.Size()
		for i := 0;i < size ; i++{
			item := queue.Dequeue().(*BFSNode)
			fmt.Println(item.val)
			newList = append(newList, item.val)
			if item.left != nil {
				queue.Enqueue(item.left)
			}
			if item.right != nil {
				queue.Enqueue(item.right)
			}
		}
	}
	return newList
}

func initBFSTree() *BFSTree {
	bfsTree := &BFSTree{}
	fmt.Println(bfsTree)
	return bfsTree
}
