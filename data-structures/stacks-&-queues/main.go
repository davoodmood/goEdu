package main

import "fmt"

// Stack reporesents -- stack that hold a slice
type Stack struct {
	items []int
}

// Stack method: Push
func (itemSlice *Stack) Push(num int) {
	itemSlice.items = append(itemSlice.items, num)
}

// Stack method: Pop
func (itemSlice *Stack) Pop() int {
	itemsLength := len(itemSlice.items) - 1
	lastOut := itemSlice.items[itemsLength]
	itemSlice.items = itemSlice.items[:itemsLength]
	
	return lastOut
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	// Push will add value at the end/top of the stack
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)

	// Pop will remove a value from the end/top of the stack
	myStack.Pop()
	fmt.Println(myStack)

	myQueue := Queue{}

	// Push will add value at the end/top of the Queue
	myQueue.Enqueue(400)
	myQueue.Enqueue(500)
	myQueue.Enqueue(600)

	fmt.Println(myQueue)

	// Pop will remove a value from the end/top of the Queue
	myQueue.Dequeue()
	fmt.Println(myQueue)
}

// TODO create a queue struct
type Queue struct {
	items []int

}

// TODO add mehtod for "enqueue"
func (itemSlice *Queue) Enqueue(num int) {
	itemSlice.items = append(itemSlice.items, num)
}

// TODO add mehtod for "dequeue"
func (itemSlice *Queue) Dequeue() int {
	lastOut := itemSlice.items[0]
	itemSlice.items = itemSlice.items[1:]

	return lastOut
}
