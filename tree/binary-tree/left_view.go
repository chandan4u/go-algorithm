package binary_tree

func (tree *BinaryTree) BinaryLeftView(level int, result map[int]int64) {
	if tree == nil {
		return
	}
	tree.Root.LeftView(level, result)
}

func (node *BinaryNode) LeftView(level int, result map[int]int64) {
	if node == nil {
		return
	}
	node.Left.LeftView(level+1, result)
	node.Right.LeftView(level+1, result)
	_, ok := result[level]
	if !ok {
		result[level] = node.Value
	}
}
