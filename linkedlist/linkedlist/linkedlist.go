package linkedlist

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
	next  *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Append adds an Item to the end of the linked list
func (ll *ItemLinkedList) Append(item int) {
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
func (ll *ItemLinkedList) Insert(i int, item int) error {
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

	return nil
}

func (ll *ItemLinkedList) Size() int {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	return ll.size
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
