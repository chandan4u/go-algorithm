package main

import "fmt"

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value int64
}

type BinaryTree struct {
	Root *BinaryNode
}

func (tree *BinaryTree) insert(value int64) *BinaryTree {
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

func (tree *BinaryTree) searchItem(item int64) {
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
	fmt.Println(">>>>>>>> Left", node.Value)
	if !ok {
		result[level] = node.Value
	}
}

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

func main() {
	// Insert Tree
	tree := &BinaryTree{}
	tree.insert(10)
	tree.insert(20)
	tree.insert(5)
	tree.insert(50)
	fmt.Printf("Result %s", tree)

	// Iteration over the tree
	staticTree := &BinaryTree{
		Root: &BinaryNode{
			Left: &BinaryNode{
				Left: nil,
				Right: &BinaryNode{
					Left: nil,
					Right: &BinaryNode{
						Left: nil,
						Right: nil,
						Value: 12,
					},
					Value: 11,
				},
				Value: 10,
			},
			Right: &BinaryNode{
				Left: nil,
				Right: &BinaryNode{
					Left: nil,
					Right: &BinaryNode{
						Left: nil,
						Right: &BinaryNode{
							Left: nil,
							Right: nil,
							Value: 90,
						},
						Value: 70,
					},
					Value: 40,
				},
				Value: 20,
			},
			Value: 15,
		},
	}
	staticTree.searchItem(100)

	// Tree left view
	resultLeft := make(map[int]int64)
	staticTree.BinaryLeftView(0, resultLeft)
	fmt.Println(resultLeft)

	resultRight := make(map[int]int64)
	staticTree.BinaryRightView(0, resultRight)
	fmt.Println(resultRight)


}
