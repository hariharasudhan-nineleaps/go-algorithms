package hashtable

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
	"github.com/hariharasudhan-nineleaps/go-algorithms/linkedlist/linkedlist"
)

// Key the key of the dictionary
type Key generic.Type

// Value the content of the dictionary
type Value generic.Type

type ValueHashtable struct {
	items map[int]*linkedlist.ItemLinkedList
	lock  sync.RWMutex
}

func New() *ValueHashtable {
	return &ValueHashtable{}
}

func (ht *ValueHashtable) Put(key Key, value Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	hash := hashcode(key)
	if ht.items == nil {
		ht.items = make(map[int]*linkedlist.ItemLinkedList)

		ll := linkedlist.ItemLinkedList{}
		ll.Append(value)

		ht.items[hash] = &ll
		return
	}

	if ht.items[hash] == nil {
		ll := linkedlist.ItemLinkedList{}
		ll.Append(value)
		ht.items[hash] = &ll

		return
	}

	ht.items[hash].Append(value)
}

func (ht *ValueHashtable) Remove(key Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	hash := hashcode(key)
	delete(ht.items, hash)
}

func (ht *ValueHashtable) Get(key Key) *linkedlist.ItemLinkedList {
	ht.lock.RLock()
	defer ht.lock.RLock()

	hash := hashcode(key)
	return ht.items[hash]
}

func (ht *ValueHashtable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RLock()

	return len(ht.items)
}

func hornerPolynomialStringHash(k Key) int {
	key := fmt.Sprintf("%s", k)
	primenumber := 31
	hash := 0

	for i := range key {
		hash += primenumber * int(key[i])
	}

	return hash
}

func hashcode(k Key) int {
	return hornerPolynomialStringHash(k)
}
