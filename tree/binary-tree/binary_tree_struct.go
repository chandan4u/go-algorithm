package binary_tree

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value int64
}

type BinaryTree struct {
	Root *BinaryNode
}
