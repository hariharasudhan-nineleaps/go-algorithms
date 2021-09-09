package main

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

func (stack *ItemStack) New() *ItemStack {
	stack.items = []Item{}
	return stack
}

func (stack *ItemStack) Push(item Item) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.items = append(stack.items, item)
}

func (stack *ItemStack) Pop() Item {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[0 : len(stack.items)-1]
	return item
}

func (stack *ItemStack) length() int {
	return len(stack.items)
}

func (stack *ItemStack) IsEmpty() bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	return stack.length() == 0
}

func (stack *ItemStack) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	return stack.length()
}

func main() {
	itemStack := ItemStack{}
	itemStack.New()

	fmt.Println("New----------------------->")
	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	fmt.Println(itemStack.items)

	fmt.Println("Push----------------------->")
	itemStack.Push(1)
	itemStack.Push(2)
	itemStack.Push(3)

	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	fmt.Println(itemStack.items)

	fmt.Println("Pop----------------------->")
	itemStack.Pop()
	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	fmt.Println(itemStack.items)

	fmt.Println("Pop----------------------->")
	itemStack.Pop()
	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	fmt.Println(itemStack.items)
}
