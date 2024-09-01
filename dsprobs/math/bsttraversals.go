package math

import "fmt"

type TraversalNode struct {
	val   int
	left  *TraversalNode
	right *TraversalNode
}

type TraversalTree struct {
	root *TraversalNode
}

func (b *TraversalTree) Insert(val int) {
	b.root = recTraversalInsert(b.root, val)
}

func recTraversalInsert(root *TraversalNode, val int) *TraversalNode {
	if root == nil {
		root = &TraversalNode{val: val}
	} else if val > root.val {
		root.right = recTraversalInsert(root.right, val)
	} else {
		root.left = recTraversalInsert(root.left, val)
	}
	return root
}

func (b *TraversalTree) Search(val int) bool {
	return recTraversalSearch(b.root, val)
}

func recTraversalSearch(root *TraversalNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.val {
		return true
	}
	if val > root.val {
		return recTraversalSearch(root.right, val)
	}
	return recTraversalSearch(root.left, val)
}

func (b TraversalTree) InOrderTraversal(output *[]int) {
	b.inOrderTraversal(b.root,output)
}

func (b TraversalTree) inOrderTraversal(node *TraversalNode, output *[]int) {
	if node != nil {
		b.inOrderTraversal(node.left,output)
		*output = append(*output, node.val)
		b.inOrderTraversal(node.right,output)
	}
}

func (b TraversalTree) PreOrderTraversal(output *[]int) {
	b.preOrderTraversal(b.root,output)
}

func (b TraversalTree) preOrderTraversal(node *TraversalNode, output *[]int) {
	if node != nil {
		*output = append(*output, node.val)
		b.preOrderTraversal(node.left, output)
		b.preOrderTraversal(node.right, output)
	}
}

func (b TraversalTree) PostOrderTraversal(output *[]int) {
	b.postOrderTraversal(b.root,output)
}


func (b TraversalTree) postOrderTraversal(node *TraversalNode, output *[]int) {
	if node != nil {
		b.postOrderTraversal(node.left,output)
		b.postOrderTraversal(node.right,output)
		*output = append(*output, node.val)
	}
}

func initTraversalTree() *TraversalTree {
	bstTree := &TraversalTree{}
	fmt.Println(bstTree)
	return bstTree
}
