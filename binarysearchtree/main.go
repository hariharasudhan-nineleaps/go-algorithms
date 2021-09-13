package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/binarysearchtree/binarysearchtree"
)

func main() {
	tree := binarysearchtree.ItemBinarySearchTree{}

	values := []int{150, 100, 90, 155, 120, 85, 200, 500, 30, 115, 125, 124, 126}
	for item := range values {
		tree.Insert(values[item], values[item])
	}

	tree.String()
	fmt.Println("Searching key 30     -->", tree.Search(30))
	fmt.Println("Searching key 130    -->", tree.Search(130))
	fmt.Println("Total count          -->", tree.Count())
	fmt.Println("Min value            -->", tree.Min())
	fmt.Println("Max value            -->", tree.Max())
	fmt.Println("Pre-Order traversal  -->", tree.PreOrderTraverse())
	fmt.Println("In-Order traversal   -->", tree.InOrderTraverse())
	fmt.Println("Post-Order traversal -->", tree.PostOrderTraverse())
	fmt.Println("Removing key 120      -->", tree.Remove(120))
	tree.String()
}
