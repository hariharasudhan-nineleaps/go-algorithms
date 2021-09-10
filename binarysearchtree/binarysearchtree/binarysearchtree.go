package binarysearchtree

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	key   int
	value Item
	left  *Node
	right *Node
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func insert(node *Node, key int, value Item) {
	if node.key < key {
		if node.left == nil {
			node.left = &Node{key, value, nil, nil}
		} else {
			insert(node.left, key, value)
		}
	} else {
		if node.right == nil {
			node.right = &Node{key, value, nil, nil}
		} else {
			insert(node.right, key, value)
		}
	}
}

func search(node *Node, key int) bool {
	if node.key == key {
		return true
	}

	if node.key < key && node.left != nil {
		return search(node.left, key)
	} else if node.key > key && node.right != nil {
		return search(node.right, key)
	}

	return false
}

func count(node *Node, counter *int) {
	*counter++

	if node.left != nil {
		count(node.left, counter)
	}

	if node.right != nil {
		count(node.right, counter)
	}
}

func (tree *ItemBinarySearchTree) Insert(key int, value Item) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	if tree.root == nil {
		tree.root = &Node{key, value, nil, nil}
		return
	}

	insert(tree.root, key, value)
}

func (tree *ItemBinarySearchTree) Search(key int) bool {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	if tree.root == nil {
		return false
	}

	return search(tree.root, key)
}

func (tree *ItemBinarySearchTree) Min() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
}

func (tree *ItemBinarySearchTree) Max() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
}

func (tree *ItemBinarySearchTree) Count() int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	counter := 0
	if tree.root == nil {
		return counter
	} else {
		count(tree.root, &counter)
	}

	return counter
}

func (tree *ItemBinarySearchTree) Remove(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
}

func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}
