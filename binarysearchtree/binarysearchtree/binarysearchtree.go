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
	if key < node.key {
		if node.left == nil {
			fmt.Println("Inserting", key, "left to node", node.key)
			node.left = &Node{key, value, nil, nil}
		} else {
			insert(node.left, key, value)
		}
	} else {
		if node.right == nil {
			fmt.Println("Inserting", key, "right to node", node.key)
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

	if key < node.key && node.left != nil {
		return search(node.left, key)
	} else if key > node.key && node.right != nil {
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

func minKey(node *Node) int {
	if node.left != nil {
		return minKey(node.left)
	}

	return node.key
}

func maxKey(node *Node) int {
	if node.right != nil {
		return maxKey(node.right)
	}

	return node.key
}

func preOrderTraverse(node *Node, values *[]int) {
	if node != nil {
		(*values) = append(*values, node.key)
		preOrderTraverse(node.left, values)
		preOrderTraverse(node.right, values)
	}
}

func inOrderTraverse(node *Node, values *[]int) {
	if node != nil {
		inOrderTraverse(node.left, values)
		(*values) = append(*values, node.key)
		inOrderTraverse(node.right, values)
	}
}

func postOrderTraverse(node *Node, values *[]int) {
	if node != nil {
		postOrderTraverse(node.left, values)
		postOrderTraverse(node.right, values)
		(*values) = append(*values, node.key)
	}
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}

	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	// Node with no children
	if node.left == nil && node.right == nil {
		node = nil
		return node
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	leftmostrightside := node.right
	for leftmostrightside.right != nil && leftmostrightside.left != nil {
		leftmostrightside = leftmostrightside.left
	}
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = remove(node.right, node.key)
	return node
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

func (tree *ItemBinarySearchTree) Min() int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	if tree.root == nil {
		return 0
	}

	return minKey(tree.root)
}

func (tree *ItemBinarySearchTree) Max() int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	if tree.root == nil {
		return 0
	}

	return maxKey(tree.root)
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

// Root, Left, Right
func (tree *ItemBinarySearchTree) PreOrderTraverse() []int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	values := []int{}
	if tree.root == nil {
		return values
	}
	preOrderTraverse(tree.root, &values)
	return values
}

// Left, Root, Right
func (tree *ItemBinarySearchTree) InOrderTraverse() []int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	values := []int{}
	if tree.root == nil {
		return values
	}

	inOrderTraverse(tree.root, &values)
	return values
}

// Left, Right, Root
func (tree *ItemBinarySearchTree) PostOrderTraverse() []int {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	values := []int{}
	if tree.root == nil {
		return values
	}

	postOrderTraverse(tree.root, &values)
	return values
}

func (tree *ItemBinarySearchTree) Remove(key int) int {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	remove(tree.root, key)
	return 1
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
