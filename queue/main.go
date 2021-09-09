package main

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

func (queue *ItemQueue) New() *ItemQueue {
	queue.items = []Item{}
	return queue
}

func (queue *ItemQueue) Enqueue(item Item) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.items = append(queue.items, item)
}

func (queue *ItemQueue) Dequeue() *Item {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	item := queue.items[0]
	queue.items = queue.items[1:len(queue.items)]

	return &item
}

func (queue *ItemQueue) Front() *Item {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	item := queue.items[0]
	return &item
}

func (queue *ItemQueue) length() int {
	return len(queue.items)
}

func (queue *ItemQueue) IsEmpty() bool {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	return queue.length() == 0
}

func (queue *ItemQueue) Size() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	return queue.length()
}

func main() {
	itemQueue := ItemQueue{}
	itemQueue.New()

	fmt.Println("Empty----------------------->")
	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	fmt.Println(itemQueue.items)

	fmt.Println("Enqueue----------------------->")
	itemQueue.Enqueue(1)
	itemQueue.Enqueue(2)
	itemQueue.Enqueue(3)

	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	fmt.Println(itemQueue.items)

	fmt.Println("Dequeue----------------------->")

	itemQueue.Dequeue()

	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	fmt.Println(itemQueue.items)

	fmt.Println("Dequeue----------------------->")

	itemQueue.Dequeue()
	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	fmt.Println(itemQueue.items)

}
