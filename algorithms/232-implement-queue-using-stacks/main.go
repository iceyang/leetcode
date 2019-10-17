package main

type Stack struct {
	nodes []int
}

func (stack *Stack) push(s int) {
	stack.nodes = append(stack.nodes, s)
}

func (stack *Stack) peek() int {
	if len(stack.nodes) == 0 {
		return 0
	}
	return stack.nodes[len(stack.nodes)-1]
}

func (stack *Stack) pop() int {
	value := stack.peek()
	stack.nodes = stack.nodes[:len(stack.nodes)-1]
	return value
}

func (stack *Stack) size() int {
	return len(stack.nodes)
}

type MyQueue struct {
	stack Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Stack{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	this.reverseStack()
	value := this.stack.pop()
	this.reverseStack()
	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	this.reverseStack()
	value := this.stack.peek()
	this.reverseStack()
	return value
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.size() == 0
}

func (this *MyQueue) reverseStack() {
	stack := Stack{}
	for this.stack.size() != 0 {
		stack.push(this.stack.pop())
	}
	this.stack = stack
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
