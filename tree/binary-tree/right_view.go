package binary_tree

func (tree *BinaryTree) BinaryRightView(level int, result map[int]int64) {
	if tree == nil {
		return
	}
	tree.Root.RightView(level, result)
}

func (node *BinaryNode) RightView(level int, result map[int]int64) {
	if node == nil {
		return
	}
	node.Left.RightView(level+1, result)
	node.Right.RightView(level+1, result)
	result[level] = node.Value
}
