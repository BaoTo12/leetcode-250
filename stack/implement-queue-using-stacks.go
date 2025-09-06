package main

type MyQueue struct {
	list []int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		list: []int{},
	}
}

func (mq *MyQueue) Push(x int) {
	mq.list = append(mq.list, x)
}

func (mq *MyQueue) Pop() int {
	if mq.Empty() {
		return -1
	}
	result := mq.list[0]

	mq.list = mq.list[1:]
	return result
}

func (mq *MyQueue) Peek() int {
	if mq.Empty() {
		return -1
	}
	return mq.list[0]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.list) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// void push(int x) Pushes element x to the back of the queue.
// int pop() Removes the element from the front of the queue and returns it.
// int peek() Returns the element at the front of the queue.
// boolean empty() Returns true if the queue is empty, false otherwise.
