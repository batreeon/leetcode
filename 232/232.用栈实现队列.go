/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
package main
type stack []int
// 	下面这三个方法都必须用指针做接收者，不然会报错，因为你需要修改接收者
func (s *stack) push(x int) {
	*s = append(*s,x)
}
func (s *stack) top() int {
	return (*s)[len(*s)-1]
}
func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

type MyQueue struct {
	in stack
	out stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	in,out := stack{},stack{}
	return MyQueue{in,out}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.in.push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var x int
	for len(this.in) > 0 {
		x = this.in.top()
		this.out.push(x)
		this.in.pop()
	}
	pop := x
	this.out.pop()
	for len(this.out) > 0 {
		x = this.out.top()
		this.in.push(x)
		this.out.pop()
	}
	return pop
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	var x int
	for len(this.in) > 0 {
		x = this.in.top()
		this.out.push(x)
		this.in.pop()
	}
	peek := x
	for len(this.out) > 0 {
		x = this.out.top()
		this.in.push(x)
		this.out.pop()
	}
	return peek
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

