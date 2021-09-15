package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/linkedlist/linkedlist"
)

func main() {
	ll := linkedlist.ItemLinkedList{}

	fmt.Println("=======Append=========")
	values := []int{150, 100, 90, 155, 120, 85, 200, 500, 30, 115, 125, 124, 126}
	for item := range values {
		ll.Append(values[item])
	}
	ll.String()
	fmt.Println("Size ====>", ll.Size())

	fmt.Println("=======Inserting=========")
	fmt.Println("Inserting item on index 3")
	ll.Insert(3, 91)
	fmt.Println("Inserting item on index 5")
	ll.Insert(5, 191)
	ll.String()
	fmt.Println("Size ====>", ll.Size())

	fmt.Println("=======Removing=========")
	fmt.Println("Removing item on index 4 ==> ")
	ll.RemoveAt(4)
	ll.String()
	fmt.Println("Size ====>", ll.Size())

	fmt.Println("=======Indexof=========")
	fmt.Println("Index of 120 ====>", ll.IndexOf(120))
	fmt.Println("Index of 2000 ====>", ll.IndexOf(2000))

	fmt.Println("=======Head=========")
	fmt.Println("LL head ====>", ll.Head())

}
