package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/queue/queue"
)

func main() {
	itemQueue := queue.ItemQueue{}
	itemQueue.New()

	fmt.Println("Empty----------------------->")
	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	itemQueue.Print()

	fmt.Println("Enqueue----------------------->")
	itemQueue.Enqueue(1)
	itemQueue.Enqueue(2)
	itemQueue.Enqueue(3)

	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	itemQueue.Print()

	fmt.Println("Dequeue----------------------->")

	itemQueue.Dequeue()

	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	itemQueue.Print()

	fmt.Println("Dequeue----------------------->")

	itemQueue.Dequeue()
	fmt.Println(itemQueue.IsEmpty())
	fmt.Println(itemQueue.Size())
	itemQueue.Print()

}
