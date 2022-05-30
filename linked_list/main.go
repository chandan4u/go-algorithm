package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Value int64
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) insert(value int64) {
	if list.Head == nil {
		list.Head = &Node{
			Next:  nil,
			Value: value,
		}
		return
	}
	list.Head.insert(value)
}

func (node *Node) insert(value int64) {
	if node.Next == nil {
		node.Next = &Node{
			Next: nil,
			Value: value,
		}
	} else {
		node.Next.insert(value)
	}
}

func (list *LinkedList) LinkedListReverse() {
	if list.Head == nil {
		return
	}
	node := list.Head
	previous := &Node{}
	for node != nil {
		temp := node.Next
		node.Next = previous
		previous = node
		node = temp
	}
	out, _ := json.Marshal(previous)
	fmt.Println(">>> Prev", string(out))
}

func main() {
	fmt.Println("Linked List")
	list := &LinkedList{}
	list.insert(10)
	list.insert(20)
	list.insert(30)
	out, _ := json.Marshal(list)
	fmt.Printf("%s", string(out))

	list.LinkedListReverse()
}
