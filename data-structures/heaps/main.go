package main

import "fmt"

// We start by creating a struct for the Max Heap

type MaxHeap struct {
	nodes []int
}

// Creating an Insert Method
func (h *MaxHeap) Insert(_key int) {
	h.nodes = append(h.nodes, _key)
	h.maxHeapifyUp(len(h.nodes)-1)
}

// Writing the Extract Method
func (h *MaxHeap) Extract() int {
	extracted := h.nodes[0]
	lastIndex := len(h.nodes) - 1

	if len(h.nodes) == 0 {
		fmt.Println("Cannot extract. There is no Node in the Tree.")
		return -1
	}

	h.nodes[0] = h.nodes[lastIndex]
	h.nodes = h.nodes[:lastIndex] // removing the last index of the slice
	h.maxHeapifyDown(0)
	return extracted
}


// @dev will heapify from bottom to tup
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.nodes[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// @dev will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.nodes) - 1
	_left, _right = left(index), right(index)
	childToCompare := 0
	for _left <= lastIndex {
		if _left == lastIndex {
			childToCompare = _left
		} else if h.nodes[_left] > h.nodes[_right] {
			childToCompare = _left
		} else {
			childToCompare = _right
		}

		if h.nodes[index] < h.nodes[ChildToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			_left, _right = left(index), right(index)
		} else {
			return
		}
	}
}


// @dev- get the parrent index
func parent (_index int) int {
	return (_index - 1) / 2
}

// @dev get the left child index
func left(_index int) int {
	return _index * 2 + 1 
}

// @dev get the right child index
func right(_index int) int {
	return _index * 2 + 2 
}

// @dev swap the keys in the slice (array)
func swap (h *MaxHeap) swap(i1, i2 int) {
	h.nodes[i1],h.nodes[i2] = h.nodes[i2],h.nodes[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10,20,30,5,8,9,11,13,15,17}


	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i:=0, i< 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}