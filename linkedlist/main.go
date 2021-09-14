package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/linkedlist/linkedlist"
)

func main() {
	ll := linkedlist.ItemLinkedList{}

	values := []int{150, 100, 90, 155, 120, 85, 200, 500, 30, 115, 125, 124, 126}
	for item := range values {
		ll.Append(values[item])
	}

	ll.String()
	fmt.Println("Size ====>", ll.Size())
	fmt.Println("Inserting item on 3 ====>", ll.Insert(3, 91))
}
