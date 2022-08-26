package main

import "fmt"
import bt "go-algorithm/tree/binary-tree"

func main() {
	// Insert Tree
	tree := &bt.BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(50)
	fmt.Printf("Result %s", tree)

	// Iteration over the tree
	staticTree := &bt.BinaryTree{
		Root: &bt.BinaryNode{
			Left: &bt.BinaryNode{
				Left: nil,
				Right: &bt.BinaryNode{
					Left: nil,
					Right: &bt.BinaryNode{
						Left: nil,
						Right: nil,
						Value: 12,
					},
					Value: 11,
				},
				Value: 10,
			},
			Right: &bt.BinaryNode{
				Left: nil,
				Right: &bt.BinaryNode{
					Left: nil,
					Right: &bt.BinaryNode{
						Left: nil,
						Right: &bt.BinaryNode{
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
	staticTree.SearchItem(100)

	// Tree left view
	resultLeft := make(map[int]int64)
	staticTree.BinaryLeftView(0, resultLeft)
	fmt.Println(resultLeft)

	// Tree right view
	resultRight := make(map[int]int64)
	staticTree.BinaryRightView(0, resultRight)
	fmt.Println(resultRight)
}
