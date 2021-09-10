package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/stack/stack"
)

func main() {
	itemStack := stack.ItemStack{}
	itemStack.New()

	fmt.Println("New----------------------->")
	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	itemStack.Print()

	fmt.Println("Push----------------------->")
	itemStack.Push(1)
	itemStack.Push(2)
	itemStack.Push(3)

	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	itemStack.Print()

	fmt.Println("Pop----------------------->")
	itemStack.Pop()
	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	itemStack.Print()

	fmt.Println("Pop----------------------->")
	itemStack.Pop()
	fmt.Println(itemStack.IsEmpty())
	fmt.Println(itemStack.Size())
	itemStack.Print()
}
