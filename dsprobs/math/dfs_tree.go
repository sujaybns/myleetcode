package math

import "fmt"

type DFSNode struct {
	val   int
	left  *DFSNode
	right *DFSNode
}

type DFSBSTTree struct {
	root *DFSNode
}

func (d *DFSBSTTree) DFSInsert(val int) {
	d.root = recDFSInsert(d.root, val)
}

func recDFSInsert(root *DFSNode, val int) *DFSNode {
	if root == nil {
		root = &DFSNode{val: val}
	} else if val < root.val {
		root.left = recDFSInsert(root.left, val)
	} else if val > root.val {
		root.right = recDFSInsert(root.right, val)
	}
	return root
}

func (d *DFSBSTTree) DFSSearch(val int) bool {
	return recDFSSearch(d.root, val)
}

func recDFSSearch(root *DFSNode, val int) bool {
	if root == nil {
		return false
	}
	if root.val == val {
		return true
	}
	if val < root.val {
		return recDFSSearch(root.left, val)
	}
	return recDFSSearch(root.right, val)
}

func (d *DFSBSTTree) DFSTraversal(result *[]int) {
	recDFSTraversal(d.root, result)
}

func recDFSTraversal(root *DFSNode, result *[]int) {
	if root == nil {
		return
	}
	fmt.Println(result)
	*result = append(*result, root.val)
	children := []*DFSNode{root.left, root.right}
	for i := 0; i < len(children); i++ {
		recDFSTraversal(children[i], result)
	}
}

func initDFSBSTTree() *DFSBSTTree {
	dfsBSTTree := &DFSBSTTree{}
	return dfsBSTTree
}
