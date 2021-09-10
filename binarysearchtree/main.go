package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/binarysearchtree/binarysearchtree"
)

func main() {
	tree := binarysearchtree.ItemBinarySearchTree{}

	values := []int{150, 100, 90, 155, 120, 85, 200, 500, 30}
	for item := range values {
		tree.Insert(values[item], values[item])
	}

	tree.String()
	fmt.Println(tree.Search(130))
	fmt.Println("Total count", tree.Count())
}
