package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Value generic.Type

type Node struct {
	value Value
	next  *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Append adds an Item to the end of the linked list
func (ll *ItemLinkedList) Append(item Value) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if ll.head == nil {
		ll.head = &Node{item, nil}
		ll.size++
		return
	}

	node := ll.head
	for node.next != nil {
		node = node.next
	}

	node.next = &Node{item, nil}
	ll.size++
}

// Insert adds an Item at position i
func (ll *ItemLinkedList) Insert(i int, item Value) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}

	if ll.head == nil || i == 0 {
		head := ll.head
		ll.head = &Node{item, head}
		ll.size++
		return nil
	}

	prevnode := ll.head
	nodeindex := 0

	for nodeindex != i-1 {
		nodeindex++
		prevnode = prevnode.next
	}

	newnode := &Node{item, prevnode.next}
	prevnode.next = newnode
	ll.size++

	return nil
}

// Removes a node at position i
func (ll *ItemLinkedList) RemoveAt(i int) (Value, error) {
	if i < 0 || i > ll.size {
		return -1, fmt.Errorf("Index out of bounds")
	}

	prevnode := ll.head
	nodeindex := 0

	for nodeindex != i-1 {
		nodeindex++
		prevnode = prevnode.next
	}

	removednode := prevnode.next
	prevnode.next = removednode.next
	ll.size--

	return removednode.value, nil
}

//Returns the position of the value
func (ll *ItemLinkedList) IndexOf(value int) int {
	if ll.head == nil {
		return -1
	}

	node := ll.head
	index := 0
	for node.next != nil && node.value != value && index != ll.size {
		node = node.next
		index++
	}

	if node.value == value {
		return index
	}

	return -1
}

//Returns the linked list size
func (ll *ItemLinkedList) Size() int {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	return ll.size
}

//Returns the first node, so we can iterate on it
func (ll *ItemLinkedList) Head() *Node {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	return ll.head
}

// Insert adds an Item at position i
func (ll *ItemLinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	j := 0
	for {
		if node == nil {
			break
		}
		j++
		fmt.Print(node.value)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}
