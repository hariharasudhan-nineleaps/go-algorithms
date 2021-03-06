package main

import (
	"fmt"

	"github.com/hariharasudhan-nineleaps/go-algorithms/hashtable/hashtable"
)

func main() {
	ht := hashtable.New()

	fmt.Println("Inserting value --> foobar into hash")
	ht.Put("foobar", "foobar")

	fmt.Println("Inserting value --> foobar into hash")
	ht.Put("foobar", "foobar")

	fmt.Println("Inserting value --> foobar into hash")
	ht.Put("johndoe", "johndoe")

	fmt.Println("Inserting value --> johndoe into hash")
	ht.Put("johndoe", "johndoe")

	fmt.Println("Inserting value --> foojohn into hash")
	ht.Put("foojohn", "johndoe")

	fmt.Println("Fetching --> foobar")
	ll1 := ht.Get("foobar")
	ll1.String()

	fmt.Println("Fetching --> foobar")
	ll2 := ht.Get("foojohn")
	ll2.String()
}
