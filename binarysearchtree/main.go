package main

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	Key   int
	Value Item
	Left  *Node
	Right *Node
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (tree *ItemBinarySearchTree) Insert(key int, value Item) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	node := &Node{key, value, nil, nil}
	if tree.root == nil {
		fmt.Println("Inserting Root Node.")
		tree.root = node
		return
	}

	fmt.Println("Root Recursive")
	insertNode(tree.root, node)
}

func insertNode(node *Node, newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left == nil {
			fmt.Println("Inserting Left Node")
			node.Left = newNode
			return
		}

		fmt.Println("Left Recursive")
		insertNode(node.Left, newNode)
		return
	} else {
		if node.Right == nil {
			fmt.Println("Inserting Right Node")
			node.Right = newNode
			return
		}

		fmt.Println("Right Recursive")
		insertNode(node.Right, newNode)
		return
	}
}

func main() {
	tree := ItemBinarySearchTree{nil, sync.RWMutex{}}
	tree.Insert(10, 10)
	tree.Insert(5, 5)
	tree.Insert(10, 10)
}
