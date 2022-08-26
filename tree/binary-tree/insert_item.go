package binary_tree

func (tree *BinaryTree) Insert(value int64) *BinaryTree {
	if tree.Root == nil {
		tree.Root = &BinaryNode{Value: value, Left: nil, Right: nil}
		return tree
	}
	tree.Root.insert(value)
	return tree
}

func (node *BinaryNode) insert(value int64) {
	if node == nil {
		return
	} else if node.Value <= value {
		if node.Left == nil {
			node.Left = &BinaryNode{Left: nil, Right: nil, Value: value}
		} else {
			node.Left.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BinaryNode{Left: nil, Right: nil, Value: value}
		} else {
			node.Right.insert(value)
		}
	}
}
