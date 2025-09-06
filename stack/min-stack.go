package main

type MinStack struct {
	list []int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		list: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.list = append(ms.list, val)
}

func (ms *MinStack) Pop() {
	ms.list = ms.list[:len(ms.list)-1]
}

func (ms *MinStack) Top() int {
	return ms.list[len(ms.list)-1]
}

func (ms *MinStack) GetMin() int {

	return 1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.
