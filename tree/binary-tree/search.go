package binary_tree

import "fmt"

func (tree *BinaryTree) SearchItem(item int64) {
	if tree.Root == nil {
		return
	}
	tree.Root.search(item)
}

func (node *BinaryNode) search(item int64) {
	if node == nil {
		return
	}
	if node.Value == item {
		fmt.Println("Found")
		return
	} else if node.Value >= item {
		node.Left.search(item)
	} else {
		node.Right.search(item)
	}
}
