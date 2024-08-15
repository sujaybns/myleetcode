package math

import "fmt"

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

type BSTTree struct {
	root *BSTNode
}

func (b *BSTTree) Insert(val int) {
	b.root = recInsert(b.root, val)
}

func recInsert(root *BSTNode, val int) *BSTNode {
	if root == nil {
		root = &BSTNode{val: val}
	} else if val > root.val {
		root.right = recInsert(root.right, val)
	} else {
		root.left = recInsert(root.left, val)
	}
	return root
}

func (b *BSTTree) Delete(val int) {
	if b.Search(val) {
		b.root = recDelete(b.root, val)
	} else {
		fmt.Println("val NOT found")
	}
}

func recDelete(root *BSTNode, val int) *BSTNode {

	if val > root.val {
		root.right = recDelete(root.right, val)
	} else if val < root.val {
		root.left = recDelete(root.left, val)
	} else {
		if root.left == nil && root.right == nil {
			return nil
		} else if root.left == nil || root.right == nil {
			if root.left == nil {
				return root.right
			}
			return root.left
		} else {
			temp := root.right
			for temp.left != nil {
				temp = temp.left
			}
			root.val = temp.val
			root.right = recDelete(root.right, root.val)
		}
	}
	return root
}

func (b *BSTTree) Search(val int) bool {
	return recSearch(b.root, val)
}

func recSearch(root *BSTNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.val {
		return true
	}
	if val > root.val {
		return recSearch(root.right, val)
	}
	return recSearch(root.left, val)
}

func (b *BSTTree) LeftSearch() *BSTNode {
	return recLeftSearch(b.root)
}

func recLeftSearch(bSTNode *BSTNode) *BSTNode {
	if bSTNode == nil {
		fmt.Printf("Reached the end")
		return nil
	}
	fmt.Printf("proceeding from :%v\n", bSTNode.val)
	bSTNode.left = recLeftSearch(bSTNode.left)
	fmt.Printf("\nbacktracking from :%v\n", bSTNode.val)
	return bSTNode
}

func (b BSTTree) inOrderTraversal(node *BSTNode) {
	if node != nil {
		b.inOrderTraversal(node.left)
		fmt.Println("\n", node.val)
		b.inOrderTraversal(node.right)
	}
}

func initBSTTree() *BSTTree {
	bstTree := &BSTTree{}
	fmt.Println(bstTree)
	return bstTree
}
