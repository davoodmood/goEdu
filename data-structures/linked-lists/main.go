package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// Defining the method happens outside the struct
// funct (Reciever struct) [method name] (arguments)
// when using * Pointer Reciever it means we want to make modifications to the reciever therefore we are accessing it directly in memory.
// if we use the Value Reciever it means we are going to work on a copy of it.
func (l *linkedList) prepend(n *node) {
	// l.head.next, l.head = l.head, n
	second := l.head
	l.head = n
	l.head.next = second
	l.length += 1
}

// In this method we use Value Reciever because we just wanna print the list and do not intend to make changes to the list and its members.
func (l linkedList) printListData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Println("")
}

func (l *linkedList) deleteNodebyValue(value int) {
	// edge cases for if the linkList is empty
	if l.length == 0 {
		return
	}

	// edge case for deleting the header
	if l.head.data == value {
		l.head = l.head.next
		l.length -= 1
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {

		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length -= 1
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 15}
	node4 := &node{data: 30}
	node5 := &node{data: 45}
	node6 := &node{data: 50}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	fmt.Printf("The Head Node is: %v", myList)
	fmt.Println("")
	myList.printListData()
	myList.deleteNodebyValue(30)
	myList.printListData()

	// checking for edge case of deleting head
	myList.deleteNodebyValue(50)
	myList.printListData()

	//checking for edge case of deleting from empty list 
	emptyList := linkedList{}
	emptyList.printListData()

	//checking for edge case of deleting non-existing node value
	myList.deleteNodebyValue(2)
	myList.printListData()

}