package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/heaps/heaps"
)

func main() {
	mh := heaps.MaxHeap{}
	mh.New()

	items := [5]int{1, 2, 3, 4, 5}

	mh.Print()
	for item := range items {
		fmt.Println("-------> Heapify Up")
		mh.Insert(items[item])
		mh.Print()
	}

	for item := range items {
		_ = item
		fmt.Println("-------> Heapify Down")
		mh.Extract()
		mh.Print()
	}

}
