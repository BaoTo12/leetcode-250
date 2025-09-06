package main

type MyStack struct {
	list []int
}

func Constructor() MyStack {
	return MyStack{
		list: []int{},
	}
}

func (ms *MyStack) Push(x int) {
	ms.list = append(ms.list, x)
}

func (ms *MyStack) Pop() int {
	result := ms.list[len(ms.list)-1]
	ms.list = ms.list[:len(ms.list)-1]
	return result
}

func (ms *MyStack) Top() int {
	return ms.list[len(ms.list)-1]
}

func (ms *MyStack) Empty() bool {
	return len(ms.list) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
