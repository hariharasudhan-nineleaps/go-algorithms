package heaps

import (
	"fmt"
	"sync"
)

type MaxHeap struct {
	items []int
	lock  sync.RWMutex
}

func (mh *MaxHeap) New() *MaxHeap {
	mh.items = make([]int, 0)
	return mh
}

func (mh *MaxHeap) Print() {
	fmt.Println(mh.items)
}

func parent(index int) int {
	return (index - 1) / 2
}

func (mh *MaxHeap) swap(i1, i2 int) {
	mh.items[i1], mh.items[i2] = mh.items[i2], mh.items[i1]
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func (mh *MaxHeap) Insert(item int) {
	mh.lock.Lock()
	defer mh.lock.Unlock()

	mh.items = append(mh.items, item)
	mh.heapifyUp(len(mh.items) - 1)
}

func (mh *MaxHeap) Extract() (extrated int) {
	mh.lock.Lock()
	defer mh.lock.Unlock()

	if len(mh.items) == 0 {
		return -1
	}

	extrated = mh.items[0]
	mh.items[0] = mh.items[len(mh.items)-1]
	mh.items = mh.items[:len(mh.items)-1]

	mh.heapifyDown(0)

	return extrated
}

func (mh *MaxHeap) heapifyUp(index int) {
	for mh.items[parent(index)] < mh.items[index] {
		mh.swap(index, parent(index))
		index = parent(index)
	}
}

func (mh *MaxHeap) heapifyDown(index int) {
	lastIndex := len(mh.items) - 1
	childToCompare := 0
	l, r := left(index), right(index)

	for l <= lastIndex { // if l is less than or equal last index ===> node has one children atleast
		if l == lastIndex {
			childToCompare = l
		} else if mh.items[l] > mh.items[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if mh.items[index] < mh.items[childToCompare] {
			mh.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}
